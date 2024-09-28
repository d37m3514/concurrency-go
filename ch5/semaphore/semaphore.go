package ch5_semaphore

import "sync"

type Semaphore struct {
	Permits int
	Cond    *sync.Cond
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		Permits: n,
		Cond:    sync.NewCond(&sync.Mutex{}),
	}
}

// Acquire() will subtract 1 to the permit when acquire is successful, then will use Wait()
// if there are no avaible permits yet.
func (s *Semaphore) Acquire() {
	s.Cond.L.Lock()
	defer s.Cond.L.Unlock()
	for s.Permits <= 0 {
		s.Cond.Wait()
	}
	s.Permits--
}

func (s *Semaphore) Release() {
	s.Cond.L.Lock()
	defer s.Cond.L.Unlock()
	s.Permits++
	s.Cond.Signal()
}
