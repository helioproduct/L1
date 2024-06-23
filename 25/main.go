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

func main() {
	start := time.Now()
	Sleep(time.Second * 5)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
