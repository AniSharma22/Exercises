package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("number of CPUs", runtime.NumCPU())
	fmt.Println("Number of Go rountines", runtime.NumGoroutine())

	counter := 100
	const gs = 50000
	var wg sync.WaitGroup
	wg.Add(gs)
	var m sync.Mutex
	for i := 0; i < gs; i++ {

		go func() {
			m.Lock()
			local := counter
			local++
			counter = local
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
