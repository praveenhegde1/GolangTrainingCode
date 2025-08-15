package main

import (
	"fmt"
)

const (
	finePerDay      = 10
	overdueInterest = 0.001
)

func calculateFine(overdueValue float64, daysOverdue int) float64 {
	fine := float64(daysOverdue*finePerDay) + (overdueValue * overdueInterest / 100)
	return fine
}

func sendFineToCustomer(fine float64) {
	fmt.Printf("Fine: INR %.2f\n", fine)
}

func main() {
	overdueValue := 1000.0
	daysOverdue := 5
	fine := calculateFine(overdueValue, daysOverdue)
	sendFineToCustomer(fine)
}
