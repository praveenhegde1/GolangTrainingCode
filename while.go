package main

import "fmt"

func main() {
	i := 1
	for i <= 5 {
		i++
		if i == 3 {
        	continue
        }
		fmt.Println(i)
		
	}
}