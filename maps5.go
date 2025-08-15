package main 

import "fmt"

//import "maps"

func main () {

m :=map[string]int{"one":1,"two":2,"three":3}
fmt.Println(m)

m2 :=make(map[string]int)

m2["blore"] =1
m2["paris"] =2
m2["tokyo"] =3

fmt.Println(m2)
x, y  :=m2["blore"]
clear(m2)
fmt.Println(x,y)
fmt.Println(m2)
}

