package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "file.xml"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Read the file content
	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// Print the file content
	fmt.Println(string(content))
}
