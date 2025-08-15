package main

import (
	"fmt"
)

type Deposit struct {
	Principal float64
	Rate      float64
	Duration  int
}

func CalculateInterest(d Deposit) (float64, error) {
	if d.Principal < 0 {
		return 0, fmt.Errorf("Principal amount cannot be negative")
	}

	interest := d.Principal * d.Rate / 100 * float64(d.Duration)
	totalAmount := d.Principal + interest

	if totalAmount < 0 {
		return 0, fmt.Errorf("Total amount cannot be negative")
	}

	return totalAmount, nil
}

func main() {
	oneYearDeposit := Deposit{Principal: 100000, Rate: 6.5, Duration: 1}
	twoYearsDeposit := Deposit{Principal: 100000, Rate: 6.75, Duration: 2}
	threeYearsDeposit := Deposit{Principal: 100000, Rate: 7, Duration: 3}
	fiveYearsDeposit := Deposit{Principal: 100000, Rate: 7.25, Duration: 5}
	highAmountDeposit := Deposit{Principal: 10000000, Rate: 7.5, Duration: 3}

	if amount, err := CalculateInterest(oneYearDeposit); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("One year deposit:", amount)
	}

	if amount, err := CalculateInterest(twoYearsDeposit); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Two years deposit:", amount)
	}

	if amount, err := CalculateInterest(threeYearsDeposit); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Three years deposit:", amount)
	}

	if amount, err := CalculateInterest(fiveYearsDeposit); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Five years deposit:", amount)
	}

	if amount, err := CalculateInterest(highAmountDeposit); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("High amount deposit:", amount)
	}
}

