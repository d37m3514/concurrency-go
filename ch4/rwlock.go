package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	mutex := sync.RWMutex{}
	var matchEvents = make([]string, 0, 10000)
	for j := 0; j < 10000; j++ {
		matchEvents = append(matchEvents, "Match event")
	}
	go matchRecorder(&matchEvents, &mutex)
	start := time.Now()
	for j := 0; j < 5000; j++ {
		go clientHandler(&matchEvents, &mutex, start)
	}
	time.Sleep(100 * time.Second)
}

func matchRecorder(matchEvents *[]string, mutex *sync.RWMutex) {
	for i := 0; ; i++ {
		mutex.Lock()
		*matchEvents = append(*matchEvents, "Match event "+strconv.Itoa(i))
		mutex.Unlock()
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Append match event")
	}
}

func clientHandler(matchEvents *[]string, mutex *sync.RWMutex, st time.Time) {
	var allEvents = make([]string, 0, 10000)
	for i := 0; i < 100; i++ {
		mutex.RLock()
		allEvents = copyAllEvents(matchEvents)
		mutex.RUnlock()
		timeTaken := time.Since(st)
		fmt.Println(len(allEvents), "events copied in", timeTaken)
	}
}

func copyAllEvents(matchEvents *[]string) []string {
	allEvents := make([]string, 0, len(*matchEvents))
	for _, e := range *matchEvents {
		allEvents = append(allEvents, e)
	}
	return allEvents
}
