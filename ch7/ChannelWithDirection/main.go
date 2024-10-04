package main

import (
	"fmt"
	"time"
)

/*
* This is to demonstrate the usage receive and send only channels.
* receive only channel (m <-chan int)
* send only channel (m chan<- int)
 */


func receiver(messages <-chan int) {
	for {
		msg := <-messages
		fmt.Println(time.Now().Format("15:04:05"), "Received: ", msg)
	}
}

func sender(messages chan<- int) {
	for i := 1; ; i++ {
		fmt.Println(time.Now().Format("15:04:05"), "Sending: ", i)
		messages <- i
		time.Sleep(1 * time.Second)
	}
}

func main() {
	msgChannel := make(chan int)
	go receiver(msgChannel)
	go sender(msgChannel)
	time.Sleep(5 * time.Second)
}
