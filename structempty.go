package main

import "fmt"

func main() {
    // Set using empty struct as value
    uniqueSet := map[string]struct{}{
        "apple":  {},
        "banana": {},
    }

    //fmt.Println("Is 'apple' in the set?", _, exists := uniqueSet["apple"]; exists)
}
