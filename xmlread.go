package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name  string `xml:"name"`
	Age   int    `xml:"age"`
	Email string `xml:"email"`
}

func main() {

	file, err := os.Open("file.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading XML file:", err)
		return
	}
	var person Person

	err = xml.Unmarshal(content, &person)
	if err != nil {
		fmt.Println("Error unmarshaling XML:", err)
		return
	}

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Email:", person.Email)
}
