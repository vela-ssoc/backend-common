package spdy

import (
	"context"
	"io"
	"math"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

type stream struct {
	id        uint32
	mux       *muxer
	syn       bool        // 是否已经发送了握手帧
	wmu       sync.Locker // 数据写锁
	cond      *sync.Cond  // stream 读写条件锁
	buff      []byte      // 消息缓冲池
	maxsize   int         // 缓冲区最大字节数
	err       error       // 错误信息
	closed    atomic.Bool // 保证 close 方法只被执行一次
	ctx       context.Context
	cancel    context.CancelFunc
	readline  time.Time
	writeline time.Time
}

func (stm *stream) Read(p []byte) (int, error) {
	return stm.read(p)
}

func (stm *stream) Write(b []byte) (int, error) {
	const max = math.MaxUint16
	bsz := len(b)
	if bsz == 0 {
		return 0, nil
	}

	stm.wmu.Lock()
	defer stm.wmu.Unlock()

	flag := flagDAT
	if !stm.syn {
		stm.syn = true
		flag = flagSYN
	}

	n := bsz
	for n > 0 {
		if n > max {
			n = max
		}

		if _, err := stm.mux.write(flag, stm.id, b[:n]); err != nil {
			return 0, err
		}

		flag = flagDAT
		b = b[n:]
		n = len(b)
	}

	return bsz, nil
}

func (stm *stream) ID() uint32           { return stm.id }
func (stm *stream) LocalAddr() net.Addr  { return stm.mux.LocalAddr() }
func (stm *stream) RemoteAddr() net.Addr { return stm.mux.RemoteAddr() }

func (stm *stream) SetDeadline(t time.Time) error {
	return stm.SetReadDeadline(t)
}

func (stm *stream) SetReadDeadline(t time.Time) error {
	return nil
}

func (stm *stream) SetWriteDeadline(time.Time) error { return nil }

func (stm *stream) Close() error {
	return stm.closeError(io.EOF, true)
}

func (stm *stream) receive(p []byte) (int, error) {
	total := len(p)
	if total == 0 {
		return 0, nil
	}

	stm.cond.L.Lock()
	for {
		psz := len(p)
		if psz == 0 {
			break
		}
		used := len(stm.buff)
		for used >= stm.maxsize {
			stm.cond.Wait()
			used = len(stm.buff)
		}

		free := stm.maxsize - used
		idx := psz
		if idx > free {
			idx = free
		}
		stm.buff = append(stm.buff, p[:idx]...)
		p = p[idx:]
	}
	stm.cond.L.Unlock()
	stm.cond.Broadcast() // 通知读取协程读取数据

	return total, nil
}

func (stm *stream) read(p []byte) (int, error) {
	psz := len(p)
	if psz == 0 {
		return 0, nil
	}

	stm.cond.L.Lock()
	used := len(stm.buff)
	// 如果缓冲区没有任何数据，就等待数据写入
	if used == 0 {
		stm.cond.Wait()
		used = len(stm.buff)
	}
	idx := psz
	if idx > used {
		idx = used
	}
	copy(p, stm.buff[:idx])
	stm.buff = stm.buff[idx:]
	stm.cond.L.Unlock()
	stm.cond.Broadcast() // 通知写入协程可以写入数据了

	return idx, nil
}

func (stm *stream) closeError(err error, fin bool) error {
	if !stm.closed.CompareAndSwap(false, true) {
		return io.ErrClosedPipe
	}

	stmID := stm.id
	stm.mux.delStream(stmID)

	if fin && stm.syn {
		_, _ = stm.mux.write(flagFIN, stmID, nil)
	}

	stm.cond.L.Lock()
	stm.err = err
	stm.cond.L.Unlock()

	stm.cancel()
	stm.cond.Broadcast()

	return err
}
