package main

import (
	"fmt"
	"sync"
)

type Job struct {
	ID  int
	Err error
}

type Worker struct {
	ID         int
	JobChannel chan Job
	Quit       chan bool
}

type WorkerPool struct {
	Workers     []*Worker
	JobChannel  chan Job
	Results     chan Job
	NumWorkers  int
	WaitGroup   sync.WaitGroup
	Quit        chan bool
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		Workers:     make([]*Worker, numWorkers),
		JobChannel:  make(chan Job),
		Results:     make(chan Job),
		NumWorkers:  numWorkers,
		WaitGroup:   sync.WaitGroup{},
		Quit:        make(chan bool),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.NumWorkers; i++ {
		wp.Workers[i] = &Worker{
			ID:         i,
			JobChannel: wp.JobChannel,
			Quit:       wp.Quit,
		}
		wp.WaitGroup.Add(1)
		go wp.Workers[i].Start(&wp.WaitGroup)
	}

	go wp.processResults()
}

func (w *Worker) Start(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case job := <-w.JobChannel:
			// Simulate some work
			if job.ID%2 == 0 {
				job.Err = fmt.Errorf("error processing job %d", job.ID)
			}
			w.processJob(job)
		case <-w.Quit:
			return
		}
	}
}

func (w *Worker) processJob(job Job) {
	// Process the job here
	fmt.Printf("Worker %d processing job %d\n", w.ID, job.ID)

	// Send the result back to the worker pool
	w.JobChannel <- job
}

func (wp *WorkerPool) processResults() {
	for {
		select {
		case job := <-wp.Results:
			if job.Err != nil {
				fmt.Printf("Error processing job %d: %s\n", job.ID, job.Err.Error())
			} else {
				fmt.Printf("Job %d processed successfully\n", job.ID)
			}
		case <-wp.Quit:
			return
		}
	}
}

func (wp *WorkerPool) AddJob(job Job) {
	wp.JobChannel <- job
}

func (wp *WorkerPool) Stop() {
	close(wp.Quit)
	wp.WaitGroup.Wait()
	close(wp.JobChannel)
	close(wp.Results)
}

func main() {
	// Create a worker pool with 5 workers
	wp := NewWorkerPool(100)
	wp.Start()

	// Add some jobs to the worker pool
	for i := 1; i < 1000; i++ {
		wp.AddJob(Job{ID: i})
	}

	// Stop the worker pool and wait for all workers to finish
	wp.Stop()
}