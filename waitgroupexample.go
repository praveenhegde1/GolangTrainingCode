package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Decrements the counter when the goroutine completes
    fmt.Printf("Worker %d starting\n", id)
    // simulate some work
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // Increment the counter
        go worker(i, &wg)
    }

    wg.Wait() // Waits for all goroutines to finish
    fmt.Println("All workers finished.")
}
