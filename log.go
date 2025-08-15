package main

import (
	"log"
	"os"
	"mathutils"
)

func main() {
	
	file, err := os.OpenFile("production.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	
	log.SetOutput(file)

	result := mathutils.Average([]float64{3, 4})
	log.Println("This is a log message", result)
	log.Printf("This is a formatted log message with a variable: %s", "example")


}