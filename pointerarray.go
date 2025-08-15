package main

import "fmt"

func modifyArray(arr *[3]int) {
    arr[0] = 100
}

func main() {
    nums := [3]int{1, 2, 3}
    modifyArray(&nums) 
    
    fmt.Println("Modified array:", nums)
}
