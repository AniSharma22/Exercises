package main

import "fmt"

type Breed int

type T interface {
	int | int8
}

type Dog struct {
	name string
	age  int
}

func (d Dog) walk() {
	fmt.Print("I am walking")
}

func (d *Dog) run() {
	fmt.Print("I am running ")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.walk()
	y.run()

}

func add[T ~int](a T, b T) {
	if func() bool {

	} {
		a + b
	}
}

func main() {
	//var a = &Dog{"kalu", 12}

	var a Breed = 10
	var b int = 20

	fmt.Print(a + b)
	//a.run()
	//fmt.Println(a)
	//youngRun(a)
}
