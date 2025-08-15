package main

import (
	"fmt"
)

func main() {
	var income float64

	for {
		fmt.Print("Enter your income (or type -1 to exit): ")
		fmt.Scan(&income)

		if income == -1 {
			break
		}

		var tax float64

		switch {
		case income <= 700000:
			tax = 0
		case income <= 1000000:
			tax = (income - 700000) * 0.10
		case income <= 1500000:
			tax = 30000 + (income-1000000)*0.15
		default:
			tax = 105000 + (income-1500000)*0.30
			surcharge := (income - 1500000) * 0.30 * 0.10 
			tax += surcharge
		}

		fmt.Printf("The calculated tax is: %.2f\n", tax)
	}
}
