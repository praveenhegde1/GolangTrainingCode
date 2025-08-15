package main

import 
(	"fmt"
    "math"
)

func main() {
	num := 16
	squareRoot := math.Sqrt(float64(num))
	fmt.Printf("The square root of %d is %.1f\n", num, squareRoot)
}