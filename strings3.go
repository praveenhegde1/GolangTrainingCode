package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Praveen Bangalore"
	reversed := reverseString(str)
	fmt.Println(reversed)
}

func reverseString(str string) string {

	chars := strings.Split(str, "")

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	reversed := strings.Join(chars, "")

	return reversed
}
