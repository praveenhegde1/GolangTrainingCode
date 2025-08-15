package main

import (
	"fmt"
)

func calculatePenaltyAndInterest(penaltyRate float64, interestRate float64) func(float64, float64) float64 {
	return func(amount, new float64)  float64 {
		penalty := amount * penaltyRate
		interest := amount * interestRate
		return penalty + interest+new
	}
}

func main() {
	
	penaltyAndInterest := calculatePenaltyAndInterest(0.05, 0.1)
	amountDue := 1000.0
	totalAmount := penaltyAndInterest(amountDue,2.0)
	fmt.Printf("Total amount due: %.2f\n", totalAmount)
}