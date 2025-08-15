package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    p := Person{Name: "Eve", Age: 35}

    // Convert struct to JSON
    jsonData, _ := json.Marshal(p)
    fmt.Println("JSON:", string(jsonData))

    // Convert JSON to struct
    var p2 Person
    json.Unmarshal(jsonData, &p2)
    fmt.Println("Struct:", p2)
}
