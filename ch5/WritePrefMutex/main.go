package main

import (
	"fmt"
	"time"
)

func main() {
	wprwmutex := NewReadWriteMutex()
	// TODO: Implementation of main

}

func readResource(cond *ReadWriteMutex, data *[]int) {
	cond.ReadLock()
	defer cond.ReadUnlock()

	for i := 0; i < len(*data); i++ {
		fmt.Printf("Reading data: %d", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Reading done.")
}

func updateLastResource(cond *ReadWriteMutex, data *[]int) {
	cond.WriteLock()
	defer cond.WriteUnlock()
	(*data)[len(*data)-1] = 100

}
