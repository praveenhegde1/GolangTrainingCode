package main

import "fmt"

func calculateIncomeTax(income float64) float64 {
	var tax float64

	if income <= 700000 {
		tax = 0
	} else if income <= 1000000 {
		tax = (income - 700000) * 0.1
	} else if income <= 1500000 {
		tax = (300000 * 0.1) + (income-1000000)*0.15
	} else {
		tax = (300000 * 0.1) + (500000 * 0.15) + (income-1500000)*0.3
	}

	return tax
}

func main() {
	income := 800000.0
	tax := calculateIncomeTax(income)
	fmt.Printf("Income: %.2f\n", income)
	fmt.Printf("Tax: %.2f\n", tax)
}
