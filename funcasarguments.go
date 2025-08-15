package main

import "fmt"

// Higher-order function
func applyOperation(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	add := func(x, y int) int { return x + y }
	multiply := func(x, y int) int { return x * y }
	division := func(x, y int) int {return x / y }
	subtract := func(x, y int) int { return x - y }

	fmt.Println("Add:", applyOperation(3, 4, add))
	fmt.Println("Multiply:", applyOperation(3, 4, multiply))
	fmt.Println("Division:", applyOperation(12, 3, division))
	fmt.Println("Subtract:", applyOperation(10, 2, subtract))
}
