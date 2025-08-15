package main 
import (
    "fmt"
	
)

func main() {

	src :=[]byte{1,2,3}
	ms := make([]byte,3)
	copy(ms, src)
	fmt.Println(ms)
	for _, b := range ms {
		fmt.Printf("value: %d\n", b)
	}
}