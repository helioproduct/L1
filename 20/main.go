package main

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Fields(s) // split по пробельным символам
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(ReverseWords("snow dog sun"))
}
