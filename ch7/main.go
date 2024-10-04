package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	msgChannel := make(chan int, 6)
	var wg sync.WaitGroup
	wg.Add(6)
	go receiver(msgChannel, &wg)
	for i := 0; i <= 5; i++ {
		msgChannel <- i
	}
	wg.Wait()
}
func receiver(messages chan int, wg *sync.WaitGroup) {
	for i := 0; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Received:", <-messages)
		wg.Done()
	}
}
