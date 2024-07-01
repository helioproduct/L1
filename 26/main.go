package main

import (
	"fmt"
	"strings"
)

// IsUnique проверяет, что все символы в строке уникальны, игнорируя регистр
func IsUnique(s string) bool {
	// Приводим строку к нижнему регистру для регистронезависимой проверки
	s = strings.ToLower(s)

	charMap := make(map[rune]bool)

	for _, char := range s {
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}

	return true
}

func main() {
	tests := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, test := range tests {
		fmt.Printf("Строка: %s, все символы уникальны: %v\n", test, IsUnique(test))
	}
}
