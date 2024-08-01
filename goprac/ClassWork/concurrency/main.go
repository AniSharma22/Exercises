package main

import (
	"fmt"
	"sync"
	"time"
)

// without mutex 14.1539 ms
// with mutex lock 14.4186 ms

func main() {

	var counter int
	const gs = 50000
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex
	startTime := time.Now()
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v = ((v + 100) * (v - 100)) % 7
			counter = v
			mu.Unlock()
			wg.Done()
		}()

	}

	wg.Wait()
	endTime := time.Now()
	fmt.Println("Time elapsed: ", endTime.Sub(startTime))
}
