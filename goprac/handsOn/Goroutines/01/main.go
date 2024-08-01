package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println(runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("hello from go routine 1")
		wg.Done()
	}()

	fmt.Println(runtime.NumGoroutine())

	go func() {
		fmt.Println("hello from go routine 2")
		wg.Done()
	}()

	fmt.Println(runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("program exits")
}
