package main

import "fmt" 

type person struct{
	name string
	age int
}

func main() { 

p := &person{name: "praveen", age: 45}
p.greet()


}

func (p *person) greet(){
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}