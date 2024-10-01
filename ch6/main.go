package main

import (
	ch6_waitgroup "d37m3514/concurrency-go/ch6/waitgroups"
	"fmt"
)

func main() {
	wg := *ch6_waitgroup.NewWaitGroupV2()
	base := 0

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go doWork(&wg, &base)
	}
	wg.Wait()
	fmt.Println("All done")
}
func doWork(wg *ch6_waitgroup.WaitGroupV2, base *int) {
	defer wg.Done()
	*base++
	fmt.Println("Done!")
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
