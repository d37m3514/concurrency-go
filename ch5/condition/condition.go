package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	money := 100
	// Creates a new mutex
	mutex := sync.Mutex{}

	// Creates a new condition variable using mutex
	cond := sync.NewCond(&mutex)

	// Passes the condition variable to both goroutines
	go stingy(&money, cond)
	go spendy(&money, cond)

	// Allow time for goroutine to execute
	time.Sleep(2 * time.Second)

	// Display the final amount of money
	mutex.Lock()
	fmt.Println("Money in the bank account: ", money)
	mutex.Unlock()
}

// Stingy() will use the Lock method to lock the shared variable and updates it,
// once the variable has been updated, it will use the Signal() method to
// notify other goroutine that the shared variable has been updated, then unlocks it.

func stingy(money *int, cond *sync.Cond) {
	for i := 0; i < 1000000; i++ {
		cond.L.Lock()
		*money += 10
		if *money >= 50 {
			cond.Signal()
		}
		cond.L.Unlock()
	}
	fmt.Println("Stingy done!")
}

// Spendy() will use the Lock() method to avoid other goroutine from accessing
// the shared variable, then check if the shared variable is less than the amount to be
// deducted, if not, it will use the Wait() method to pause the execution atomically and unlocks
// the execution to let other goroutine gain access

func spendy(money *int, cond *sync.Cond) {
	for i := 0; i < 200000; i++ {
		cond.L.Lock()
		for *money < 50 {
			cond.Wait()
		}
		if *money <= 0 {
			fmt.Println("Money is negative!")
			os.Exit(1)
		}
		*money -= 50
		cond.L.Unlock()
	}
	fmt.Println("Spendy done!")
}
