package main

import (
	"fmt"
	"time"
)

/*
This is to demonstrate the usage receive and send only channels.
receive only channel (m <-chan int)
send only channel (m chan<- int)

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
*/

/*
A channel can also be closed instead of using a sentinel value, or "poison pill" signal like -1 to close a channel.
We can use the built-in close() method to close a channel when needed
*/

func main() {
	msgChan := make(chan int)
	go receiver(msgChan)
	for i := 1; i < 3; i++ {
		fmt.Println(time.Now().Format("15:04:05"), "Sending:", i)
		msgChan <- i
		time.Sleep(1 * time.Second)
	}
	close(msgChan)
	time.Sleep(5 * time.Second)
}
func receiver(message <-chan int) {
	for {
		msg, more := <-message
		if !more {
			fmt.Println(time.Now().Format("15:04:05"), "Channel closed:", msg, more)
			time.Sleep(1 * time.Second)
		} else {
			fmt.Println(time.Now().Format("15:04:05"), "received:", msg, more)
			time.Sleep(1 * time.Second)
		}
	}
}
