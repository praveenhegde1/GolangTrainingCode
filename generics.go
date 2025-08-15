package main

import "fmt"


func Swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {

	x := 10
	 y := 20

	x, y = Swap(x,y)
	fmt.Println("Swapped integers:", x, y)

	
	str1, str2 := "Hello", "World"
	str1, str2 = Swap(str1, str2)
	fmt.Println("Swapped strings:", str1, str2)
}