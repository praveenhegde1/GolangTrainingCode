package main

import "fmt"

// Define nested structs
type Address struct {
    City  string
    State string
}

type Person struct {
    Name    string
    Age     int
    Address Address
}

func main() {
    p := Person{
        Name: "Charlie",
        Age:  40,
        Address: Address{
            City:  "New York",
            State: "NY",
        },
    }

    fmt.Println("Name:", p.Name)
    fmt.Println("City:", p.Address.City)
}
