package main

import "fmt"


func main() {

	
		arr := make([][]int, 5)
		for i := 0; i < 5; i++ {
			arr[i] = make([]int, 3)
		}
	fmt.Println(arr)
		// Accessing and modifying elements
		arr[0][0] = 1
		arr[1][2] = 3
	
		// Printing the array
		fmt.Println(arr)
	}
	





