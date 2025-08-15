package main

import "fmt"

func main() {
	names := map[string]string{
		"pradeep": "pradeep",
		"jane":    "Jane Smith",
		"praveen": "Praveen Kumar",
		"alice":   "Alice Johnson",
	}

	for key, value := range names {
		if key == "praveen" {
			fmt.Println("Found name:", value)
			break
		}
	}
}
