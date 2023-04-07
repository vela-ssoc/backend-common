package queue

import (
	"context"
	"sync"
	"sync/atomic"
)

type Runner interface {
	Run()
}

type Tasker interface {
	Start()
	Submit(Runner)
	Shutdown()
}

func NewTask(workers, buff int) Tasker {
	if workers <= 0 || workers > 100 {
		workers = 8
	}
	if buff <= 0 || buff > 1024 {
		buff = 1024
	}

	tasks := make(chan Runner, buff)
	ctx, cancel := context.WithCancel(context.Background())

	return &queue{
		tasks:   tasks,
		workers: workers,
		ctx:     ctx,
		cancel:  cancel,
	}
}

type queue struct {
	tasks   chan Runner
	workers int
	working atomic.Bool
	wg      sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
}

func (q *queue) Submit(rn Runner) {
	if rn == nil {
		return
	}
	if err := q.ctx.Err(); err == nil {
		q.tasks <- rn
	}
}

func (q *queue) Shutdown() {
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
		case rn := <-q.tasks:
			q.call(rn)
		}
	}
}

func (*queue) call(rn Runner) {
	defer func() { recover() }()
	rn.Run()
}
