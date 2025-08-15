package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World!"

	if strings.HasPrefix(str, "Hello") {
		fmt.Println("String starts with 'Hello'")
	}

	if strings.HasSuffix(str, "World!") {
		fmt.Println("String ends with 'World!'")
	}

	index := strings.Index(str, "World")
	if index != -1 {
		fmt.Println("Substring 'World' found at index", index)
	}

	newStr := strings.ReplaceAll(str, "Hello", "Hi")
	fmt.Println("New string:", newStr)

	parts := strings.Split(str, ",")
	fmt.Println("Parts:", parts)

	joinedStr := strings.Join(parts, "-")
	fmt.Println("Joined string:", joinedStr)

	upperStr := strings.ToUpper(str)
	fmt.Println("Uppercase string:", upperStr)

	lowerStr := strings.ToLower(str)
	fmt.Println("Lowercase string:", lowerStr)
}
