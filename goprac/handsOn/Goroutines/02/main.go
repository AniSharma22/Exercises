package main

import (
	"fmt"
)

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Println("Hello, playground")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{"anish"}
	saySomething(&p1)

	// saySomething(p1)  does not work
}
