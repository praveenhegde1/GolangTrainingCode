package main

import "fmt"

func main() {
	// Creating a dynamic 2D array
	var rows, cols int
	fmt.Print("Enter the number of rows: ")
	fmt.Scan(&rows)
	fmt.Print("Enter the number of columns: ")
	fmt.Scan(&cols)

	arr := make([][]int, rows)
	for i := 0; i < rows; i++ {
		arr[i] = make([]int, cols)
	}

	// Accessing and modifying elements
	arr[0][0] = 1
	arr[1][2] = 3

	// Printing the array
	fmt.Println(arr)
}
