package ch6_waitgroup

import "sync"

type Barrier struct {
	groupSize int
	cond      *sync.Cond
}

func NewBarrier() *Barrier {
	return &Barrier{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}
func (wg *Barrier) Add(delta int) {
	wg.cond.L.Lock()
	wg.groupSize += delta
	wg.cond.L.Unlock()
}
func (wg *Barrier) Wait() {
	wg.cond.L.Lock()
	for wg.groupSize > 0 {
		// Waits and atomically releases the mutex while groupSize is greater than 0
		wg.cond.Wait()
	}
	wg.cond.L.Unlock()
}

func (wg *Barrier) Done() {
	wg.cond.L.Lock()
	defer wg.cond.L.Unlock()
	wg.groupSize--
	if wg.groupSize == 0 {
		wg.cond.Broadcast()
	}
}
