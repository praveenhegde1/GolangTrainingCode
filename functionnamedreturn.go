package main

import "fmt"

// Function with named return values
func rectangleProperties(length, width float64) (area, perimeter,volume float64) {
	area = length * width
	perimeter = 2 * (length + width)
	volume = length * width * length // Assuming length, width, and height are the same for volume calculation
	return // Named return allows omission of variables here
}

func main() {
	area, perimeter,volume := rectangleProperties(5.0, 3.0)
	fmt.Println("Area:", area, "Perimeter:", perimeter, "Volume:", volume)
}
