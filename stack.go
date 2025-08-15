package main
import "fmt"
func greet() *string{
	name := "kolkota"
	return &name
}

func main(){
	greet()
}
