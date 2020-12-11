package main

import (
	"fmt"
)

func main() {
	var num = 100
	var p = &num

	fmt.Println("num before = ", num)

	// Changing the value stored in the pointed variable through the pointer
	*p = 5

	fmt.Println("num after = ", num)
}
