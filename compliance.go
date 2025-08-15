package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type ComplianceCheck struct {
	Regulation string
	Status     string
	Duedate    time.Time
}

func Reminder(check *ComplianceCheck) {
	time.Sleep(time.Until(check.Duedate))
	fmt.Printf("Reminder: %s is due\n", check.Regulation)
}

func main() {
	checks := []*ComplianceCheck{
		{Regulation: "Environmental Regulation", Status: "Due", Duedate: time.Now().Add(1 * time.Second)},
		{Regulation: "Health Regulation", Status: "Due", Duedate: time.Now().Add(24 * time.Hour)},
		{Regulation: "Safety Regulation", Status: "Due", Duedate: time.Now().Add(48 * time.Hour)},
		{Regulation: "Security Regulation", Status: "Due", Duedate: time.Now().Add(72 * time.Hour)},
		{Regulation: "Tax Regulation", Status: "Overdue", Duedate: time.Now().Add(-48 * time.Hour)},
	}

	var wg sync.WaitGroup
	runtime.SetBlockProfileRate(1)
	fmt.Println("Starting goroutines...", runtime.NumGoroutine())


	runtime.GC()	
	fmt.Println("GC stats: Done")
	fmt.Println(runtime.MemStats{})
	fmt.Println(runtime.NumCPU())

	for _, check := range checks {
		wg.Add(1)
		go func(check *ComplianceCheck) {
			defer wg.Done()
			fmt.Printf("Processing %s check\n", check.Regulation)
			Reminder(check)
		}(check)
		fmt.Println("# of goroutines after creating:...", runtime.NumGoroutine())
		wg.Wait()
		fmt.Println("# of gorouitnes after wait group:-- ", runtime.NumGoroutine())
	}
}

