package main

import "fmt"
import "log"

func Swap(a, b *int64) {
	*a += *b
	*b = *a - *b
	*a -= *b
}

func main() {

	var a, b int64

	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Scanf("%d", &b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("a = %d, b = %d\n", a, b)
	Swap(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b)
}
