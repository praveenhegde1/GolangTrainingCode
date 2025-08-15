package main 
import (
	"fmt"
)

func main() {
	creditCardNumber := "1234567890123456" 

	if len(creditCardNumber) < 16 {
		fmt.Println("Invalid credit card number")
		return
	}

	maskedNumber :=  creditCardNumber[12:]
	fmt.Println("Masked Credit Card Number:", maskedNumber)
}