package main

import "fmt"

type person struct {
	name string
	age  int
	email string
}
func main() {

	person := person{
		name : "praveen",
		age :  45,
		email : "praveen@saastech.com",
    }
	fmt.Println(person)
	fmt.Println(person.name)
	fmt.Println(person.age)



}