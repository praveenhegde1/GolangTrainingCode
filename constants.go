package main

import "fmt"

const (
	OptionA = 1  << iota
	OptionB 
	OptionC
	OptionD
)

func main() {
	
	fmt.Println(OptionA) // Output: 1
	fmt.Println(OptionB) // Output: 2
	fmt.Println(OptionC) // Output: 4
	fmt.Println(OptionD) // Output: 8
}