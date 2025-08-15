package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "100" }()

	msg := <-messages
	fmt.Println(msg)
}
