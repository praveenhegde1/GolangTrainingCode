package main

import "fmt"

func main() {
	// Static array with predefined size
	var staticArray [5]int
	staticArray[0] = 10
	staticArray[1] = 20
	staticArray[2] = 30
	staticArray[3] = 40
	staticArray[4] = 50
	
	// Dynamic array using slice
	dynamicArray := []int{1, 2, 3, 4, 5,6,8,9,10,11,12}

	// Array iteration using for loop
	for i := 0; i < len(staticArray); i++ {
		fmt.Println(staticArray[i])
	}

	// Array iteration using range
	for _, value := range dynamicArray {
		fmt.Println(value)
	}
}