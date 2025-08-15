package main

import (
	"fmt"
)

type Stack[T any] []T

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Pop() (T, error) {
	/*if len(*s) == 0 {
		
		return 0, fmt.Errorf("stack is empty")
	} */
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value, nil
}

func main() {
	stack := Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	for {
		value, err := stack.Pop()
		if err != nil {
			break
		}
		fmt.Println(value)
	}
}

