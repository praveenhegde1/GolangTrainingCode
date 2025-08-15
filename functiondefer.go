package main

import "fmt"

// Function demonstrating defer
func deferredFunction() {
	defer fmt.Println("This is executed last")
	fmt.Println("This is executed first")
}

func main() {
	deferredFunction()
}
