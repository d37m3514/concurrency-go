package main

import (
	"sync"
)

// readersCounter stores the number of readers currently holding the read lock
// writersWaiting stores the number of writers currently waiting
// writerActive indicates if writer is holding a write lock
type ReadWriteMutex struct {
	readersCounter int
	writersWaiting int
	writerActive   bool
	cond           *sync.Cond
}

// NewReadWriteMutex() initialize a new ReadWriteMutex with a new condition variable
// and associated mutex.
func NewReadWriteMutex() *ReadWriteMutex {
	return &ReadWriteMutex{cond: sync.NewCond(&sync.Mutex{})}
}

// ReadLock() will hold the read lock by using Wait() then wait before there are no more
// waiting or active writer. Once the conditions are met, increment the readersCounter by 1
// then release the lock
func (rw *ReadWriteMutex) ReadLock() {
	rw.cond.L.Lock()
	defer rw.cond.L.Unlock()
	for rw.writersWaiting > 0 || rw.writerActive {
		rw.cond.Wait()
	}
	rw.readersCounter++
}

// WriteLock() will acquire the mutex to have exclusive execution
// mutex and condition variable is used to wait as long as readers or writers
// are active. In addition, it increments the writersWaiting variable to indicate that
// it's waiting for the lock to become available. Once the writer's lock has been acquired,
// decrement the writer waiting by 1, and set the writerActive to true.
func (rw *ReadWriteMutex) WriteLock() {
	rw.cond.L.Lock()
	defer rw.cond.L.Unlock()
	rw.writersWaiting++
	for rw.readersCounter > 0 || rw.writerActive {
		rw.cond.Wait()
	}
	rw.writersWaiting--
	rw.writerActive = true
}

// ReadUnlock() will decrement the readers counter, and if the reader counter is 0,
// it means it is the last goroutine and will use Broadcast() to notify reader and writer
// goroutine that are waiting
func (rw *ReadWriteMutex) ReadUnlock() {
	rw.cond.L.Lock()
	defer rw.cond.L.Unlock()
	rw.readersCounter--
	if rw.readersCounter == 0 {
		rw.cond.Broadcast()
	}
}
