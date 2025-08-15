package main

import "fmt"

func main() {
	var num int = 10
	var ptr *int

	ptr = &num

	fmt.Println("Value of num before modification:", num)
	
	*ptr = 20

	fmt.Println("Value of num after modification:", num)
}