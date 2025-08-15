package main

import (
	"fmt"
	"time"
)

func main() {
	
	rateLimit := time.Tick(time.Second)

	
	for i := 0; i < 10; i++ {
		<-rateLimit 
		go makeRequest(i)
	}

	
	time.Sleep(time.Second * 10)
}

func makeRequest(requestNum int) {
	fmt.Printf("Making request %d\n", requestNum)
	// Perform the actual request here
}