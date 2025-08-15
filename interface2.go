package main

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) string
	Refund(amount float64) string
}

type CreditCard struct {
	CardNumber string
}

func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card %s", amount, cc.CardNumber)
}

func (cc CreditCard) Refund(amount float64) string {
	return fmt.Sprintf("Refunded %.2f to Credit Card %s", amount, cc.CardNumber)
}

type PayPal struct {
	Email string
}

func (pp PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal account %s", amount, pp.Email)
}

func (pp PayPal) Refund(amount float64) string {
	return fmt.Sprintf("Refunded %.2f to PayPal account %s", amount, pp.Email)
}

func ProcessPayment(pm PaymentMethod, amount float64) {
	fmt.Println(pm.Pay(amount))
	fmt.Println(pm.Refund(amount))
}

func main() {
	cc := CreditCard{CardNumber: "1234-5678-9012-3456"}
	pp := PayPal{Email: "user@example.com"}

	fmt.Println("Processing Credit Card Payment:")
	ProcessPayment(cc, 100.0)

	fmt.Println("\nProcessing PayPal Payment:")
	ProcessPayment(pp, 200.0)
}
