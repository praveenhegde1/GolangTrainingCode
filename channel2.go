package main

import "fmt"

func main() {

	messages := make(chan string,5)
	messages <- "Hello"
	messages <- "World"
	messages <- "!"	
	messages <- "Bye"
	messages <- "Text"
	
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages) 

}
