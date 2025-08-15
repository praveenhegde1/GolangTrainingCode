package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(2 * time.Second)

	go func() {

		time.Sleep(3 * time.Second)

		if !timer.Stop() {
			<-timer.C
		}

		fmt.Println("Task completed")
	}()

	<-timer.C

	fmt.Println("Timer expired")
}
