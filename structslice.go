package main

import "fmt"

type Student struct {
    Name  string
    Grade int
}

func main() {
    students := []Student{
        {Name: "Alice", Grade: 90},
        {Name: "Bob", Grade: 85},
    }

    for _, s := range students {
        fmt.Println("Name:", s.Name, "Grade:", s.Grade)
    }
}
