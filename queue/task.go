package queue

import (
	"context"
	"sync"
	"sync/atomic"
)

type Tasker interface {
	Start()
	Add(func())
	Stop()
}

func NewTask(workers, buff int) Tasker {
	if workers <= 0 || workers > 100 {
		workers = 8
	}
	if buff <= 0 || buff > 1024 {
		buff = 1024
	}

	tasks := make(chan func(), buff)
	ctx, cancel := context.WithCancel(context.Background())

	return &queue{
		tasks:   tasks,
		workers: workers,
		ctx:     ctx,
		cancel:  cancel,
	}
}

type queue struct {
	tasks   chan func()
	workers int
	working atomic.Bool
	wg      sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
}

func (q *queue) Add(fn func()) {
	if err := q.ctx.Err(); err == nil {
		q.tasks <- fn
	}
}

func (q *queue) Stop() {
	if q.working.CompareAndSwap(true, false) {
		q.cancel()
		q.wg.Wait()
	}
}

func (q *queue) Start() {
	if !q.working.CompareAndSwap(false, true) {
		return
	}

	n := q.workers
	q.wg.Add(n)
	for i := 0; i < n; i++ {
		go q.worker()
	}
}

func (q *queue) worker() {
	defer q.wg.Done()

over:
	for {
		select {
		case <-q.ctx.Done():
			break over
		case fn := <-q.tasks:
			if fn == nil {
				continue
			}
			q.call(fn)
		}
	}
}

func (*queue) call(fn func()) {
	defer func() { recover() }()
	fn()
}
