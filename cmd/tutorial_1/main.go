package main

import (
	"fmt"
)

func main() {
	intSlice := []int32{1, 2, 3}
	fmt.Printf("our array is: %v", intSlice)

	intSlice = append(intSlice, 4)
	fmt.Printf("our array after slicing: %v", intSlice)
}
