package ch6_waitgroup

import ch5_semaphore "d37m3514/concurrency-go/ch5/semaphore"

type WaitGroup struct {
	sema *ch5_semaphore.Semaphore
}

func NewWaitGroup(size int) *WaitGroup {
	return &WaitGroup{sema: ch5_semaphore.NewSemaphore(1 - size)}
}

func (wg *WaitGroup) Wait() {
	wg.sema.Acquire()
}

func (wg *WaitGroup) Done() {
	wg.sema.Release()
}
