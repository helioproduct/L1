package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(arr []int) []int {
	// Если массив пустой или состоит из одного элемента, он уже отсортирован
	if len(arr) < 2 {
		return arr
	}

	// Выбор опорного элемента
	pivotIndex := rand.Intn(len(arr))
	pivot := arr[pivotIndex]

	// Разделение массива на три части: меньшие, равные и большие опорному элементу
	left, right := 0, len(arr)-1
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	for i := range arr {
		if arr[i] < pivot {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	// Рекурсивная сортировка левой и правой частей
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])

	return arr
}

func main() {
	fmt.Println("nil:", QuickSort(nil))
	fmt.Println("len = 1:", QuickSort([]int{5}))
	fmt.Println("len > 1:", QuickSort([]int{5, 3, 7, 6, 5, 4, 4, 6, 2, 8, 7}))
	fmt.Println("reversed:", QuickSort([]int{6, 5, 4, 3, 2, 1}))
	fmt.Println("sorted:", QuickSort([]int{1, 2, 3, 4, 5, 6}))
}
