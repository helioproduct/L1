package main

import (
	"fmt"
	"sync"
	// "time"
)

type Counter struct {
	Value int
	taken chan bool
}

func NewCounter() *Counter {
	counter := new(Counter)
	counter.taken = make(chan bool, 1)
	counter.Value = 0
	return counter
}

func (c *Counter) Increment() {
	c.taken <- true
	c.Value += 1
	<-c.taken
}

func main() {
	var wg sync.WaitGroup
	counter := NewCounter()

	wg.Add(4)
	go func() {
		// time.Sleep(time.Second * 1)
		counter.Increment()
		wg.Done()
	}()
	go func() {
		counter.Increment()
		wg.Done()
	}()
	go func() {
		counter.Increment()
		wg.Done()
	}()
	go func() {
		counter.Increment()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(counter.Value)
}
