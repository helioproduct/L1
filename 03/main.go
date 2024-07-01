package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var sum atomic.Uint64

	sourceArray := []uint64{2, 4, 6, 8, 10}

	// Запускаем горутину для каждого числа в массиве
	for _, number := range sourceArray {
		wg.Add(1)
		go func(wg *sync.WaitGroup, delta uint64) {
			sum.Add(delta)
			wg.Done()
		}(&wg, number*number)
	}

	wg.Wait()
	fmt.Println(sum.Load())
}
