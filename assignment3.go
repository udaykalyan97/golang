package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scanln(&a)
	fmt.Println(strings.Title(strings.ToLower(a)))

}
