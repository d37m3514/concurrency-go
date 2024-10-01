package ch6_waitgroup

import "sync"

type WaitGroupV2 struct {
	groupSize int
	cond      *sync.Cond
}

func NewWaitGroupV2() *WaitGroupV2 {
	return &WaitGroupV2{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}
func (wg *WaitGroupV2) Add(delta int) {
	wg.cond.L.Lock()
	wg.groupSize += delta
	wg.cond.L.Unlock()
}
func (wg *WaitGroupV2) Wait() {
	wg.cond.L.Lock()
	for wg.groupSize > 0 {
		// Waits and atomically releases the mutex while groupSize is greater than 0
		wg.cond.Wait()
	}
	wg.cond.L.Unlock()
}

func (wg *WaitGroupV2) Done() {
	wg.cond.L.Lock()
	defer wg.cond.L.Unlock()
	wg.groupSize--
	if wg.groupSize == 0 {
		wg.cond.Broadcast()
	}
}
