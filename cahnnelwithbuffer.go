package main

import (
	"fmt"
	"time"
)

func main() {
	// Buffered channel with a capacity of 2
	ch := make(chan int, 2)

	// Start 5 tasks
	for i := 1; i <= 5; i++ {
		ch <- i // Send task to channel (blocks if full)
		go func(task int) {
			fmt.Printf("Processing task %d\n", task)
			time.Sleep(1 * time.Second) // Simulate work
			<-ch // Remove task from buffer
		}(i)
	}

	// Wait for all tasks to complete
	time.Sleep(6 * time.Second)
}
