package main

import "fmt"
import "time"

// https://ru.hexlet.io/qna/go/questions/golang-kak-ostanovit-gorutinu#:~:text=%D0%95%D1%81%D0%BB%D0%B8%20%D0%BA%D0%B0%D0%BD%D0%B0%D0%BB%20%D0%BF%D0%BE%D0%BB%D1%83%D1%87%D0%B0%D0%B5%D1%82%20%D1%81%D0%B8%D0%B3%D0%BD%D0%B0%D0%BB%20%D0%BE%D0%B1,%D0%B7%D0%B0%D0%B2%D0%B5%D1%80%D1%88%D0%B5%D0%BD%D0%B8%D1%8F%20%D1%80%D0%B0%D0%B1%D0%BE%D1%82%D1%8B%20%D0%B3%D0%BE%D1%80%D1%83%D1%82%D0%B8%D0%BD%D1%8B%2C%20%D0%BF%D1%80%D0%BE%D0%B3%D1%80%D0%B0%D0%BC%D0%BC%D0%B0%20%D0%B7%D0%B0%D0%B2%D0%B5%D1%80%D1%88%D0%B0%D0%B5%D1%82%D1%81%D1%8F.
// https://metanit.com/go/tutorial/7.3.php
func main() {

	channel := make(chan string)
	stop := make(chan bool)

	go func() {
		for {
			select {
			default:
				channel <- "data"
			case <-stop:
				close(channel)
				return
			}
		}
	}()

	go func() {
		for {
			select {
			default:
				fmt.Println(<-channel)
			case <-stop:
				return
			}
		}
	}()

	time.Sleep(1 * time.Second)
	stop <- true
}
