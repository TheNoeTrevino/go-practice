package main

import (
	"errors"
	"fmt"
)

func main() {
	number1 := 5
	number2 := 0
	result, remainder, err := intDivision(number1, number2)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("the result is %v, and the remainder is %v", result, remainder)
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("the denominator cannot be 0")
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err
}
