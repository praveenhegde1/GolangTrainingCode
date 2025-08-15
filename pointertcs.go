package main

import "fmt"

func changvalue(v *int) {
	*v = 20
    fmt.Println("value after", *v)
}
func main() {

 x :=10
 fmt.Println("value bfore", x)

 changvalue(&x)
 fmt.Println("value after", x)

}