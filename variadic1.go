package main

import (
	"fmt"
)

type InterestStatement struct {
	PAN         string
	Year        int
	Interest    float64
	Description string
}

func ConsolidateInterestStatements(pan string, statements ...InterestStatement) float64 {
	totalInterest := 0.0

	for _, statement := range statements {
		if statement.PAN == pan {
			totalInterest += statement.Interest
		}
	}

	return totalInterest
}

func main() {
	
	statement1 := InterestStatement{PAN: "ABCD1234E", Year: 2021, Interest: 1000.0, Description: "Bank Interest"}
	statement2 := InterestStatement{PAN: "ABCD1234E", Year: 2021, Interest: 500.0, Description: "Stock Dividends"}
	statement3 := InterestStatement{PAN: "XYZW5678F", Year: 2021, Interest: 2000.0, Description: "Bank Interest"}

	
	totalInterest := ConsolidateInterestStatements("ABCD1234E", statement1, statement2, statement3)

	fmt.Printf("Total interest for PAN ABCD1234E: %.2f\n", totalInterest)
}