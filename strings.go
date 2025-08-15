package main

import (
	"fmt"
	"strings"
)

func main() {

	message := "Hello, World!"

	fmt.Println(message)

	if strings.Contains(message, "World") {
		fmt.Println("The string contains 'World'")
	} else {
		fmt.Println("The string does not contain 'World'")
	}

	words := strings.Split(message, " ")
	fmt.Println("The words in the string are:", words)

	newMessage := strings.Join(words, "-")
	fmt.Println("The new message is:", newMessage)

	replacedMessage := strings.Replace(message, "Hello", "Hi", 1)
	fmt.Println("The replaced message is:", replacedMessage)
	
}
