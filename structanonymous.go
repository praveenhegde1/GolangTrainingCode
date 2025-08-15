package main

import "fmt"

func main() {
    // Anonymous struct
    p := struct {
        Name string
        Age  int
    }{
        Name: "Bob",
        Age: 30,
    }

    fmt.Println("Name:", p.Name, "Age:", p.Age)
}
