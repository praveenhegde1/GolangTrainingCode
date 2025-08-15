package main

import "fmt"

func Print[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

func main() {

	ints := []int{1, 2, 3}
strings := []string{"a", "b", "c"}

Print(ints)
Print(strings)

}
