package main

import (
	"fmt"
	"sync"
)

// Функция для вычисления квадрата числа
func CalculateSquare(wg *sync.WaitGroup, input chan int, output chan int) {
	defer wg.Done()
	defer close(output) // Закрываем канал output после завершения функции

	for number := range input {
		output <- number * number
	}
}

// Функция для заполнения канала числами из массива
func FillFromSlice(wg *sync.WaitGroup, sourceArray *[]int, inputChannel chan int) {
	defer wg.Done()
	defer close(inputChannel) // Закрываем канал input после завершения функции

	for _, number := range *sourceArray {
		inputChannel <- number
	}
}

// Функция для чтения данных из канала и вывода их в stdout
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
	go CalculateSquare(&wg, input, output)     // Запускаем горутину для вычисления квадрата числа
	go FillFromSlice(&wg, &sourceArray, input) // Запускаем горутину для заполнения input числами из массива
	go ExhaustChannel(&wg, output)             // Запускаем горутину для чтения из output и вывода на экран

	wg.Wait()
}
