package main

import (
	"fmt"
	"runtime"
)

type Internship struct {
	Duration int
	Stipened int
	Role     string
}

func (i Internship) String() string {
	return fmt.Sprintf("hello this is default")
}

func swap(i int, j int, arr [5]int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func main() {
	var a1 Internship

	a1.Stipened = 300

	fmt.Println(a1)

	fmt.Println(runtime.NumCPU())

	a2 := [5]int{1, 2, 3, 4, 5}
	swap(2, 3, a2)
	fmt.Println(a2)

}
