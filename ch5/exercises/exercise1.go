package main

import (
	"d37m3514/concurrency-go/ch5/broadcast"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	wg := sync.WaitGroup{}
	playerPool := make([]int, 0, 10)
	playersInGame := 10
	wg.Add(2)
	for playerId := 1; playerId <= 10; playerId++ {
		go broadcast.PlayerHandler(cond, &playersInGame, playerId, &playerPool)
		time.Sleep(1 * time.Second)
	}
	go broadcast.StartInFive(cond, playerPool)
	wg.Wait()
}
