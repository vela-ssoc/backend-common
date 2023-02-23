package syscmd

func limited(max int) *limitedBuffer {
	return &limitedBuffer{
		max: max,
		buf: make([]byte, max),
	}
}

// limitedBuffer 有限 buffer，超出 buffer 的数据则会丢失
type limitedBuffer struct {
	max int    // 最大容量
	idx int    // 当前存储空间占用
	buf []byte // 存储空间
}

func (lb *limitedBuffer) Reset() {
	lb.idx = 0
}

func (lb *limitedBuffer) Write(p []byte) (int, error) {
	psz := len(p)
	sur := lb.max - lb.idx
	if n := sur; n > 0 {
		if sur > psz {
			n = psz
		}
		copy(lb.buf[lb.idx:], p[:n])
		lb.idx += n
	}
	return psz, nil
}

func (lb *limitedBuffer) bytes() []byte {
	return lb.buf[:lb.idx]
}
