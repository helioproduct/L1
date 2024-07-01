package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("1048577", 10) // 2^20 + 1
	b.SetString("2097153", 10) // 2^21 + 1

	mul := new(big.Int).Mul(a, b)
	fmt.Printf("%s * %s = %s\n", a.String(), b.String(), mul.String())

	div := new(big.Int).Div(a, b)
	fmt.Printf("%s / %s = %s\n", a.String(), b.String(), div.String())

	sum := new(big.Int).Add(a, b)
	fmt.Printf("%s + %s = %s\n", a.String(), b.String(), sum.String())

	sub := new(big.Int).Sub(a, b)
	fmt.Printf("%s - %s = %s\n", a.String(), b.String(), sub.String())
}
