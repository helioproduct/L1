package main

import (
	"fmt"
	"sync"
)

func CalculateSquare(wg *sync.WaitGroup, input chan int, output chan int) {
	defer wg.Done()
	defer close(output)

	for number := range input {
		output <- number * number
	}
}

func FillFromSlice(wg *sync.WaitGroup, sourceArray *[]int, inputChannel chan int) {
	defer wg.Done()
	defer close(inputChannel)

	for _, number := range *sourceArray {
		inputChannel <- number
	}
}

func ExhaustChannel(wg *sync.WaitGroup, sourceChannel chan int) {
	defer wg.Done()
	for number := range sourceChannel {
		fmt.Println(number)
	}
}

func main() {

	var wg sync.WaitGroup

	input := make(chan int)
	output := make(chan int)

	var sourceArray []int
	for i := 1; i < 1000; i++ {
		sourceArray = append(sourceArray, i)
	}

	wg.Add(3)
	go CalculateSquare(&wg, input, output)
	go FillFromSlice(&wg, &sourceArray, input)
	go ExhaustChannel(&wg, output)

	wg.Wait()
}
