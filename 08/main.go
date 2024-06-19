package main

import "fmt"
import "log"

func ChangeBit(number *int64, index int, bitValue int) error {
	if index < 0 || index >= 63 {
		return fmt.Errorf("index out of range")
	}
	if bitValue < 0 || bitValue > 1 {
		return fmt.Errorf("bit value should be 1 or 0")
	}
	*number &= ^(1 << index)
	*number = *number | int64(bitValue)<<index
	return nil
}

func main() {

	var number int64
	number = 230492394

	err := ChangeBit(&number, 3, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(number)
}
