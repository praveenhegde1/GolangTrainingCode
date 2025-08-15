package main

import (
	"fmt"
	
)

func calculateCompoundInterest(principal, rate, time float64) float64 {
	if time == 0 {
		return principal
	}
	interest := principal * rate / 365
	amount := principal + interest
	return calculateCompoundInterest(amount, rate, time-1)
}

func main() {
	principal := 1000.0
	rate := 8.0
	time := 45.0

	compoundInterest := calculateCompoundInterest(principal, rate, time)
	fmt.Printf("Compound Interest at the end of the month: %.2f\n", compoundInterest-principal)
}