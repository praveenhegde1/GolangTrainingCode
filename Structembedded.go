package main

import "fmt"

type Animal struct {
    Name string
}

type Dog struct {
    Animal
    Breed string
}

func main() {
    d := Dog{
        Animal: Animal{Name: "Buddy"},
        Breed:  "Golden Retriever",
    }

    // Access embedded fields
    fmt.Println("Name:", d.Name) // No need for d.Animal.Name
    fmt.Println("Breed:", d.Breed)
}
