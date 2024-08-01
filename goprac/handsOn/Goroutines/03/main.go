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
	const gs = 500
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {

		go func() {
			local := counter
			local++
			counter = local
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}
