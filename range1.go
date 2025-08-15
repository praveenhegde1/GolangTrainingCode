package main

import "fmt"

func main() {
	contributions := []int{5000, 25000, 150000, 500000, 2000000}

	for _, contribution := range contributions {
		switch {
		case contribution > 1000000:
			fmt.Println("A")
		case contribution > 100000:
			fmt.Println("B")
	case	 contribution > 10000:
			fmt.Println("C")
		default:
			fmt.Println("D")
		}
	}
}
