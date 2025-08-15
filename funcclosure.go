package main

import "fmt"

// Function that returns another function
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	inc := counter()
	fmt.Println(inc()) // 1
	fmt.Println(inc()) // 2
	fmt.Println(inc()) // 3
	fmt.Println(inc())
	
}
