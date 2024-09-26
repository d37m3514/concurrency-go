package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// creates a new condition variable
	cond := sync.NewCond(&sync.Mutex{})
	// Initialize the total number of players
	playersInGame := 4
	for playerId := 1; playerId < 5; playerId++ {
		// Starts a goroutine sharing a condition variable,
		// players in game, and player ID
		go playerHandler(cond, &playersInGame, playerId)
		// Sleeps for 1-second interval before the next player
		time.Sleep(1 * time.Second)
	}
}
func playerHandler(cond *sync.Cond, pr *int, pid int) {
	// Locks the mutex in the condition variable to avoid race condition
	cond.L.Lock()
	fmt.Println(pid, ": Connected")
	// Subtracts 1 from the shared remaining players variable
	*pr--
	// if remaining players is equal to zero, it means all players are connected,
	// and it will send a Broadcast() to all goroutine that are waiting
	if *pr == 0 {
		cond.Broadcast()
	}
	// Waits on a condition variable as long as there are more players to connect
	for *pr > 0 {
		fmt.Println(pid, ": Waiting for more players")
		cond.Wait()

	}
	// Unlocks the mutex so that all goroutine can resume execution and start the game
	cond.L.Unlock()
	fmt.Printf("Player %d is ready\n", pid)
}
