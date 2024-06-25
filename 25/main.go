package main

import "time"
import "fmt"

func Sleep(d time.Duration) {
	start := time.Now()
	for {
		if time.Now().Sub(start) >= d {
			break
		}
	}
}

func SleepWithTimer(d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
}

func main() {
	c := make(chan bool)

	start := time.Now()

	go func() {
		// time.Sleep(time.Second * 5)
		// Sleep(time.Second * 5)
		SleepWithTimer(time.Second * 5)
		c <- true
	}()

	go func() {
		// time.Sleep(time.Second * 5)
		// Sleep(time.Second * 5)
		SleepWithTimer(time.Second * 5)
		c <- true
	}()

	<-c
	<-c
	end := time.Now()
	fmt.Println(end.Sub(start))
}
