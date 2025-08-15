package main

import "fmt"

func main() {

	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	sum := 0

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {

			sum += arr[i][j]
		}
	}

	fmt.Println("Sum of array elements:", sum)
}
