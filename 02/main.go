package main

import (
	"fmt"
	"sync"
)

// PrintSquare принимает WaitGroup и число, вычисляет квадрат числа и выводит его на экран.
// После завершения работы вызывает Done() для уменьшения счетчика WaitGroup.
func PrintSquare(wg *sync.WaitGroup, number int) {
	defer wg.Done()
	fmt.Println(number * number)
}

func main() {
	var wg sync.WaitGroup
	sourceArray := []int{2, 4, 6, 8, 10}

	for _, number := range sourceArray {
		wg.Add(1)
		go PrintSquare(&wg, number)
	}

	wg.Wait()
}
