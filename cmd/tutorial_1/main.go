package main

import (
	"fmt"
)

func main() {
	intSlice := []int32{1, 2, 3}
	fmt.Printf("our array is: %v\n", intSlice)

	intSlice = append(intSlice, 4)
	fmt.Printf("our array after slicing: %v\n", intSlice)

	myMap := map[string]uint8{"noe": 24}
	myMap["josh"] = 4
	fmt.Printf("Map: %v\n", myMap)
	counter := 0
	for name, age := range myMap {
		counter += 1
		fmt.Printf("ID: %v, Name: %v, Age: %v\n", counter, name, age)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("hello")
	}
}

func print() {
	fmt.Println("hello")
}

const noe = value
