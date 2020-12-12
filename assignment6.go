package main

import (
	"fmt"
	"strconv"
)

func isNumeric(s string) bool {
	_, err := strconv.ParseInt(s, 2, 2)
	return err == nil
}

func main() {
	// Example - 1
	str := "1Ax3 7y Bk"
	digitcount := 0
	spacecount := 0
	lowercount := 0
	uppercount := 0

	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		fmt.Println(int(runes[i]))

		if 47 < int(runes[i]) && int(runes[i]) < 58 {
			digitcount += 1
		}
		if int(runes[i]) == 32 {
			spacecount += 1
		}
		if 64 < int(runes[i]) && int(runes[i]) < 91 {
			uppercount += 1
		}
		if 96 < int(runes[i]) && int(runes[i]) < 123 {
			lowercount += 1
		}
	}
	fmt.Printf("no.of digits = %d\n", digitcount)
	fmt.Printf("no.of spaces = %d\n", spacecount)
	fmt.Printf("no.of lowercase letters = %d\n", lowercount)
	fmt.Printf("no.of uppercase letters = %d\n", uppercount)

}
