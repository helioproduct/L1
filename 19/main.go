package main

import (
	"fmt"
)

// ReverseString переворачивает строку, поддерживая символы Unicode.
func ReverseString(s string) string {
	// Преобразуем строку в срез рун для корректной работы с Unicode
	runes := []rune(s)
	left, right := 0, len(runes)-1

	// Обмениваем символы местами, двигаясь от краев к центру
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	// Преобразуем срез рун обратно в строку
	return string(runes)
}

func main() {
	// Пример использования функции
	input := "главрыба"
	output := ReverseString(input)
	fmt.Printf("Исходная строка: %s\nПеревернутая строка: %s\n", input, output)
}
