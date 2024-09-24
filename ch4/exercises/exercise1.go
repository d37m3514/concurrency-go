package main

import (
	"fmt"
	"sync"
	"time"
)

func countdown(seconds *int, mutex *sync.RWMutex) {
	for *seconds > 0 {
		time.Sleep(1 * time.Second)
		mutex.Lock()
		*seconds -= 1
		mutex.Unlock()
	}
}

func main() {
	count := 5
	mutex := sync.RWMutex{}
	go countdown(&count, &mutex)

	for count > 0 {
		mutex.Lock()
		fmt.Println(count)
		mutex.Unlock()
	}
}
