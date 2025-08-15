package main

import (
	
	"fmt"
	"errors"
)

func divide(a, b int) (int,error)  {
	if (b == 0)  {
		return 0, errors.New("division by zero")
	} 
	return a / b, nil
}

func main() {
	//result, err := divide(10, 5)
	result, err := divide(10,2)
	if err != nil {
		panic(err)
		//fmt.Println("Error:", err)
		//return
	}
	
	fmt.Println("Result:", result)
}
