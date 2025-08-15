package main

import "fmt"

func main1() {
	// Logical AND
	fmt.Println(true && true)   // true
	fmt.Println(true && false)  // false
	fmt.Println(false && true)  // false
	fmt.Println(false && false) // false

	// Logical OR
	fmt.Println(true || true)   // true
	fmt.Println(true || false)  // true
	fmt.Println(false || true)  // true
	fmt.Println(false || false) // false

	// Logical NOT
	fmt.Println(!true)  // false
	fmt.Println(!false) // true
}
