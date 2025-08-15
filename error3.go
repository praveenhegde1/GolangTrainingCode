package main

import (
	"errors"
	"fmt"
)

func Add[T any](a, b T) (T, error) {
	switch a := a.(type) {
	case int:
		return a.(int) + b.(int), nil
		return a + b.(int), nil
	case float32:
		return a + b.(float32), nil
	case float64:
		return a + b.(float64), nil
	default:
		return nil, errors.New("unsupported data type")
	}
}

func main() {
	result, err := Add(10, 20)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = Add(3.14, 2.71)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	result, err = Add("Hello", "World")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}