package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)


type WorkerPool struct {
	workers int
	ch      chan http.ResponseWriter
}


func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		workers: workers,
		ch:      make(chan http.ResponseWriter, workers),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		go func() {
			for w := range wp.ch {
				fmt.Fprint(w, "Hello, World!")
				runtime.Goexit()
			}
		}()
	}
}


func (wp *WorkerPool) HandleRequest(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println(w)
	w.Header().Set("Content-Type", "text/plain")
	wp.ch <- w
}

func main() {
	wp := NewWorkerPool(100)
	wp.Start()
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err!= nil {
        log.Fatal( err)
    }

	defer file.Close()
	

	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		wp.HandleRequest(w, r)
		stack := make([]byte, 1024*1024) 
		n := runtime.Stack(stack, true)

		//fmt.Fprintf(os.Stdout,"Goroutine %d stack trace:\n%s",runtime.NumGoroutine(),string(stack[:n]))
		log.SetOutput(file)
		log.Printf("Goroutine %d stack trace:\n%s",runtime.NumGoroutine(),string(stack[:n]))
		log.Printf(r.URL.Path)
		log.SetOutput(os.Stdout)

	//	fmt.Print("Goroutine %d stack trace:\n%s",runtime.NumGoroutine(),string(stack[:n])) 
		
	})
	http.ListenAndServe(":8081", nil)
}
