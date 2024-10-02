package main

import (
	ch6_waitgroup "d37m3514/concurrency-go/ch6/waitgroups"
	"fmt"
	"time"
)

func main() {
	Barrier := *ch6_waitgroup.NewBarrier()
	go workAndWait("Red", 4, &Barrier)
	go workAndWait("Green", 2, &Barrier)
	time.Sleep(10 * time.Second)
}

func workAndWait(name string, timeToWork int, barrier *ch6_waitgroup.Barrier) {
	start := time.Now()
	fmt.Println(time.Since(start), name, "is running...")
	time.Sleep(time.Duration(timeToWork) * time.Second)
	fmt.Println(time.Since(start), name, "is waiting on barrier")
	barrier.Wait()
}

/*func main() {
	wg := *ch6_waitgroup.NewWaitGroup(4)
	start := time.Now()
	for i := 1; i <= 4; i++ {
		go doWork(i, &wg)
	}
	wg.Wait()
	fmt.Println("All complete")
	duration := time.Since(start).Truncate(time.Millisecond)
	fmt.Printf("Execution w/ goroutines lasted %v seconds", duration)
}

func doWork(id int, wg *ch6_waitgroup.WaitGroup) {
	i := rand.Intn(5)
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Printf("%d : Done working after %d seconds\n", id, i)
	wg.Done()
}*/
