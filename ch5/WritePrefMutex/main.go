package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wprwmutex := NewReadWriteMutex()
	sharedResourceSlice := []int{1, 2, 3}
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			readResource(wprwmutex, &sharedResourceSlice)
		}()
		go func() {
			defer wg.Done()
			updateLastResource(wprwmutex, &sharedResourceSlice)
		}()
		time.Sleep(2 * time.Second)
	}
	wg.Wait()
}

func readResource(cond *ReadWriteMutex, data *[]int) {
	cond.ReadLock()
	defer cond.ReadUnlock()

	for i := 0; i < len(*data); i++ {
		fmt.Printf("Reading data: %d\n", (*data)[i])
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Reading done.")
}

func updateLastResource(cond *ReadWriteMutex, data *[]int) {
	cond.WriteLock()
	defer cond.WriteUnlock()
	(*data)[len(*data)-1] = 100
	fmt.Printf("Updated data: %v\n", *data) // Print the updated slice
	fmt.Println("Write finished")
}
