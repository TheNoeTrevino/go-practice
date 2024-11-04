package main

import (
	"fmt"
)

func main() {
	var myNum int32 = 520
	intArr := [3]int32{1, 2, myNum}
	fmt.Printf("our array is: %v", intArr)
}
