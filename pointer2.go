package main

import "fmt"


type Account struct {
	balance float64
}


func NewAccount(initialBalance float64) *Account {
	return &Account{balance: initialBalance}
}


func (a *Account) Deposit(amount float64) {
	a.balance += amount
}


func (a *Account) Withdraw(amount float64) error {
	if a.balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	a.balance -= amount
	return nil
}

func main() {
	
	account := NewAccount(30)
	account.Deposit(40)
	fmt.Println("Balance:", account.balance) 

	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Balance:", account.balance) 
	}
}
