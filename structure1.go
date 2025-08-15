package main

import "fmt"

type Stud struct {
	Name  string
	Age   int
	Grade string
}

func main() {

	student := Stud{
		Name:  "Praveen Hegde",
		Age:   45,
		Grade: "A",
	}
	fmt.Println("Student Name:", student.Name)
	fmt.Println("Student Age:", student.Age)
	fmt.Println("Student Grade:", student.Grade)
}
