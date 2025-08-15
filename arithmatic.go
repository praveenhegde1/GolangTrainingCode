package main

import "fmt"

func main() {
	// Arithmetic operations for normal variables
	a := 10
	b := 5

	sum := a + b
	diff := a - b
	product := a * b
	quotient := a / b
	remainder := a % b

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", diff)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)

	// Arithmetic operations for complex variables
	var c complex64
	var d complex64
	c = complex(3, 4)
	d = complex(1, 2)

	sumComplex := c + d
	diffComplex := c - d
	productComplex := c * d
	quotientComplex := c / d

	fmt.Println("Sum (Complex):", sumComplex)
	fmt.Println("Difference (Complex):", diffComplex)
	fmt.Println("Product (Complex):", productComplex)
	fmt.Println("Quotient (Complex):", quotientComplex)

	// Additional operations for complex variables
	//conjugate := cmplx.Conj(complex128(c))
	//abs := cmplx.Abs(complex128(c))

	//	fmt.Println("Conjugate:", conjugate)
	//	fmt.Println("Absolute value:", abs)

}
