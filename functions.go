package main

import (
	"fmt"
)

func calculateBankCharges(cost float64) float64 {
	if cost <= 2000 {
		return 50
	} else if cost <= 5000 {
		return cost * 0.01
	} else if cost <= 150000 {
		return cost * 0.005
	} else {
		return 0
	}
}

func main() {
	var cost float64
	fmt.Print("Enter the cost: ")
	fmt.Scanln(&cost)

	bankCharges := calculateBankCharges(cost)
	fmt.Printf("Bank charges: %.2f\n", bankCharges)
}
