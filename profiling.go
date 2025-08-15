package main

import (
    "fmt"
    "math/rand"
    "net/http"
    _ "net/http/pprof"
    "time"
)

func work() {
    data := make([]int, 0)
    for i := 0; i < 1000000; i++ {
        data = append(data, rand.Intn(100))
    }
    time.Sleep(2 * time.Second)
}

func main() {
    go func() {
        fmt.Println("Starting pprof server on :6060")
        http.ListenAndServe("localhost:6060", nil) 
    }()
    
    for i := 0; i < 100000; i++ {
        work()
    }
}
