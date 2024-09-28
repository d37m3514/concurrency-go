package broadcast

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func PlayerHandler(cond *sync.Cond, pr *int, pid int, pp *[]int) {
	// Locks the mutex in the condition variable to avoid race condition
	cond.L.Lock()
	sleepBeforeConn(3)
	*pp = append(*pp, pid)
	fmt.Println("Player ", pid, ": Connected")
	// Subtracts 1 from the shared remaining players variable
	*pr--
	// if remaining players is equal to zero, it means all players are connected,
	// and it will send a Broadcast() to all goroutine that are waiting
	if *pr == 0 {
		cond.Broadcast()
	}
	// Unlocks the mutex so that all goroutine can resume execution and start the game
	cond.L.Unlock()
}

func StartInFive(cond *sync.Cond, players []int) {
	time.Sleep(2 * time.Second)
	cond.L.Lock()
	defer cond.L.Unlock()

	for i := 2; i >= 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Printf("Game will start in %d\n", i)
	}
	connectedPlayers(players)
}

func connectedPlayers(players []int) {
	fmt.Println("Connected players:")
	for _, v := range players {
		fmt.Println("Player ", v)
	}
}

func sleepBeforeConn(sec int) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomInt := r.Intn(sec)
	time.Sleep(time.Duration(randomInt) * time.Second)
}
