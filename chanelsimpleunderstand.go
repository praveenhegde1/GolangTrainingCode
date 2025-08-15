package main

import (
	"fmt"
)

func worker(id int, ch chan string) {
	ch <- fmt.Sprintf("Worker %d completed", id) 
}
func main() {
	ch := make(chan string)

	// Launch multiple workers
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	// Receive messages from workers
	for i := 1; i <= 3; i++ {
		fmt.Println(<-ch) // Receive message from channel
	}
}
