package main

import (
	"fmt"
	"sync"
	"time"
)

func function(d time.Duration) {

	stop := time.After(d)
	start := time.Now()

	for {
		select {
		case <-stop:
			fmt.Println("stopped afer", time.Since(start))
			return

		default:
			fmt.Println("working")
		}
	}

}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		function(3 * time.Second)
	}()
	wg.Wait()

}
