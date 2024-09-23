package rwmutex

import (
	"sync"
)

type ReadWriteMutex struct {
	// Used to count the number of reader goroutines currently in the critical section
	readersCounter int
	// Mutex for synchronizing readers access
	readersLock sync.Mutex
	// Mutex for blocking any writer access
	globalLock sync.Mutex
}

func (rw *ReadWriteMutex) ReadLock() {
	// Synchronize access so that only one goroutine is allowed at any time.
	rw.readersLock.Lock()
	// Reader goroutine increments readerCounter by 1
	rw.readersCounter++
	// If a reader goroutine is the first one in, it attempts to lock globalLock.
	if rw.readersCounter == 1 {
		rw.globalLock.Lock()
	}
	// Synchronizes access so that only one goroutine is allowed at any time.
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) ReadUnlock() {
	// Synchronizes access so that only one goroutine is allowe any time.
	rw.readersLock.Lock()
	// The reader goroutine decrements readersCounter by 1.
	rw.readersCounter--
  // If the reader goroutine is the last one out,
  // it unlocks the global lock.
	if rw.readersCounter == 0 {
		rw.globalUnlock.Unlock()
	}
	rw.readersLock.Unlock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
  rw.globalLock.Unlock()
}
