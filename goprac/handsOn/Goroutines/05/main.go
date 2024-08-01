package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func add[T int | string](a, b T) {
	fmt.Println("hello")
}

func main() {
	fmt.Println("number of CPUs", runtime.NumCPU())
	fmt.Println("Number of Go rountines", runtime.NumGoroutine())

	var counter int64
	const gs = 50000
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {

		go func() {
			atomic.AddInt64(&counter, 10)
			fmt.Println(atomic.LoadInt64(&counter))

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter", counter)
}
