package main

import "fmt"

func Search(data *[]int, target int) int {

	left := 0
	right := len(*data) - 1
	mid := int((left + right) / 2)

	for {
		midValue := (*data)[mid]
		if midValue == target {
			return mid
		} else if midValue > target {
			right = mid - 1
		} else {
			left = mid + 1
			right = len(*data) - 1
		}

		mid = int((left + right) / 2)
		if left > right {
			return -1
		}
	}
}

func main() {

	data := []int{1, 2, 3, 4, 5, 6, 15, 120, 121, 494}
	target := 494
	result := Search(&data, target)

	if result != -1 {
		fmt.Println(target, "at index of ", result)
	} else {
		fmt.Println("Не найдено")
	}
}
