package main

import "fmt"
import "mathfunction"


func main() {
	num1 := 5.0
    num2 := 3.0

    result := mathfunction.Add(num1, num2)
    fmt.Printf("Addition: %.2f + %.2f = %.2f\n", num1, num2, result)

}