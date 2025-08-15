package main

import "fmt"

func main() {
	var x uint8
	var y uint8
	x = 10
	y = 10

	if x > y {
		fmt.Println("x is greater than y")

		if x%2 == 0 {
			fmt.Println("x is even")
		} else {
			fmt.Println("x is odd")
		}
	} else if x < y {
		fmt.Println("x is less than y")

		if y%2 == 0 {
			fmt.Println("y is even")
		} else {
			fmt.Println("y is odd")
		}
	} else {
		fmt.Println("x is equal to y")
	}
}
