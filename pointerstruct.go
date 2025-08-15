package main

import "fmt"

type Person struct {
    name string
    age  int
}

func main() {
    person := Person{"Alice", 25}
    ptr := &person 
    
    fmt.Println("Name:", ptr.name) 
    fmt.Println("Age:", ptr.age)
    
    ptr.age = 30 
    fmt.Println("Updated Age:", person.age)
}
