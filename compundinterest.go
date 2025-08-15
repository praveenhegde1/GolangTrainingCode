package main

import (
	"fmt"
	"math"
)

func main() {
	principal := 1000.0
	rate := 0.1
	time := 2.0

	sendDailyNotice(principal, rate, time)
}

func calculateCompoundInterest(principal, rate, time float64) float64 {
	n := 365.0 // Compounded daily
	amount := principal * math.Pow(1+(rate/n), n*time)
	return amount - principal
}

func sendDailyNotice(principal, rate, time float64) {
	interest := calculateCompoundInterest(principal, rate, time)
	fmt.Printf("Daily Compound Interest Notice:\n")
	fmt.Printf("Principal: INR%.2f\n", principal)
	fmt.Printf("Rate: %.2f%%\n", rate)
	fmt.Printf("Time: %.2f years\n", time)
	fmt.Printf("Interest: INR%.2f\n", interest)
}

