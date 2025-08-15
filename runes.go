package main

import "fmt"

func main() {
  
    s := "Hello, World!"

   
    for i, r := range s {
        fmt.Printf("Rune %d: %c (code point: %d)\n", i, r, r)
    }
}