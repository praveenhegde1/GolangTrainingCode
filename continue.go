package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		// Skip numbers divisible by 3
		if i%3 == 0 {
			fmt.Printf("Skipping %d\n", i)
			continue
		}
		fmt.Printf("Processing %d\n", i)
	}
}
