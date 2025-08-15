package main

import "fmt"

func calculateIncomeTax(income float32) float32 {
	var tax float32

	switch {
	case income <= 700000:
		tax = 0
	case income <= 1000000:
		tax = (income - 700000) * 0.1
	case income <= 1500000:
		tax = (income-1000000)*0.15 + 30000
	default:
		tax = (income-1500000)*0.3 + 105000
	}

	if income > 1500000 {
		tax += tax * 0.05 
	}

	return tax
}

func main() {
	var income float32
	income = 3500000.0
	tax := calculateIncomeTax(income)
	fmt.Printf("Income: %.2f\nTax: %.2f\n", income, tax)
}
