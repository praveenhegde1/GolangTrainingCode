package main

import "fmt"

func main() {
	var num uint8
	num = 100

	if num > 0 {
		fmt.Println("Number is positive")
	} else if num < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is zero")
	}
}
