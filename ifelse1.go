package main

import "fmt"

func main() {
	var interestRate float32
	var principal float32
	var time float32
	interestRate = 0.05
	principal = 1000.0
	time = 1.0

	interest := principal * interestRate * time

	if interest > 0 {
		fmt.Printf("Interest earned: INR%.2f\n", interest)
	} else {
		fmt.Println("No interest earned")
	}
}
