package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

func Worker(wg *sync.WaitGroup, source chan interface{}, id int) {
	defer wg.Done()
	for element := range source {
		fmt.Printf("[WORKER %d] recieved: %s\n", id, element)
	}
}

func main() {

	fmt.Println("input workers count:")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	line = strings.TrimSpace(line)

	workersCount, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal(err)
	}

	source := make(chan interface{}, 1)

	sigs := make(chan os.Signal)
	done := make(chan bool)

	signal.Notify(sigs, syscall.SIGINT)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-done:
				close(source)
				return
			default:
				source <- "some data"
			}
		}
	}()

	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go Worker(&wg, source, i)
	}

	<-sigs
	done <- true

	wg.Wait()
}
