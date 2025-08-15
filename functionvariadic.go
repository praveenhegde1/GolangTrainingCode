package main

import "fmt"

// Function to calculate the sum of variable number of arguments
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	fmt.Println("Sum:", sum(1, 2, 3, 4, 5)) // Pass multiple arguments
	fmt.Println("Sum:", sum())             // Pass no arguments
}
