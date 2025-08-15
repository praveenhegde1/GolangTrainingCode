// interest/calculator.go
package interest

import (
    "math"
)

// CompoundInterest calculates compound interest amount
// Formula: A = P * (1 + r/n)^(nt)
// Returns Total Amount (A), Interest earned (A - P)
func CompoundInterest(principal, rate float64, timesCompoundedPerYear, years float64) (total, interest float64) {
    amount := principal * math.Pow((1 + rate/(100*timesCompoundedPerYear)), timesCompoundedPerYear*years)
    return amount, amount - principal
}
