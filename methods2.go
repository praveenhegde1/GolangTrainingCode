package main

import (
	"fmt"
	"math/rand"
	
)

type Customer struct {
	ID   int
	Name string
}

func (c *Customer) IssueToken() string {
	token := fmt.Sprintf("%s-%d", c.Name, rand.Intn(1000))
	return token
}

func main() {
	customer := Customer{
		ID:   1,
		Name: "Praveen Hegde",
	}

	token := customer.IssueToken()
	fmt.Println("Token:", token)
}
