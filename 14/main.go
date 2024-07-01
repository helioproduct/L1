package main

import (
	"fmt"
)

// determineType принимает  переменную типа interface{} и определяет её тип
func determineType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("Тип переменной: int")
	case string:
		fmt.Println("Тип переменной: string")
	case bool:
		fmt.Println("Тип переменной: bool")
	case chan int:
		fmt.Println("Тип переменной: channel")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	// Примеры переменных разных типов
	var i int = 23
	var s string = "g"
	var b bool = true
	var ch chan int = make(chan int)

	determineType(i)
	determineType(s)
	determineType(b)
	determineType(ch)
}
