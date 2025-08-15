package main

import "fmt"

func main() {
	// Anonymous function assigned to a variable
	multiply := func(a, b int) int {
		return a * b
	}

	result := multiply(4, 5)
	fmt.Println("Product:", result)
}
