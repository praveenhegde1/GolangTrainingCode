package main

import(
	"fmt"
	"errors"
)

type FinancialReport struct {
	Income  float64
	Expenses float64
	NetoperatingIncome float64

}

func GenerateReport(income,expenses float64)(*FinancialReport,error){

if income <0 || expenses <0 {
	return nil, errors.New("Income and expenses should be non-negative")
}
netOperatingIncome := income - expenses

report := &FinancialReport{Income: income, Expenses: expenses, NetoperatingIncome: netOperatingIncome,  }
return report, nil

}
    

func main() {

	income := 250000.0
	expenses := 220000.0
	report, err := GenerateReport(income, expenses)
	if err != nil {
		fmt.Println("Error:", err)
        return
    }

	fmt.Printf("Income: $%.2f\n", report.Income)
	fmt.Printf("Expenses: $%.2f\n", report.Expenses)
	fmt.Printf("Net Operating Income: $%.2f\n", report.NetoperatingIncome)
}



}
