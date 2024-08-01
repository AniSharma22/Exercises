package main

import (
	"fmt"
	"sync"
	"time"
)

// total elapsed time: 15.2822ms

func main() {
	const (
		outerGoRoutines = 500
		innerGoRoutines = 50
	)
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Add to WaitGroup for outer goroutines
	wg.Add(outerGoRoutines)

	startTime := time.Now()

	for i := 0; i < outerGoRoutines; i++ {
		go func() {
			// Decrement outer WaitGroup when done
			defer wg.Done()

			// Add to WaitGroup for inner goroutines

			for j := 0; j < innerGoRoutines; j++ {
				wg.Add(1)
				go func() {
					// Decrement inner WaitGroup when done
					defer wg.Done()
					// Safely increment the counter
					mu.Lock()
					counter++
					mu.Unlock()
				}()

			}

			// Wait for all inner goroutines to complete
		}()
	}

	// Wait for all outer goroutines to complete
	wg.Wait()
	endTime := time.Now()
	fmt.Println("total elapsed time:", endTime.Sub(startTime))

	fmt.Printf("Final counter value: %d\n", counter)
	fmt.Printf("Expected counter value: %d\n", outerGoRoutines*innerGoRoutines)
}
