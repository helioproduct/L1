package main

import (
	"fmt"
)

func removeElement(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный слайс:", slice)

	indexToRemove := 2
	newSlice := removeElement(slice, indexToRemove)
	fmt.Printf("Слайс после удаления элемента с индексом %d: %v\n", indexToRemove, newSlice)
}
