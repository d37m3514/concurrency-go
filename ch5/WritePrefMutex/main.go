package main

import (
	"fmt"
	"sync"
)

func main() {
	wprwmutex := NewReadWriteMutex()
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			readResource(wprwmutex)
		}()
		go func() {
			defer wg.Done()
			updateLastResource(wprwmutex)
		}()
	}
	wg.Wait()
}

func readResource(cond *ReadWriteMutex) {
	cond.ReadLock()
	defer cond.ReadUnlock()
	fmt.Println("Reading done.")
}

func updateLastResource(cond *ReadWriteMutex) {
	cond.WriteLock()
	defer cond.WriteUnlock()
	fmt.Println("Write finished")
}
