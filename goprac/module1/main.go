package main

import "fmt"

//go:generate stringer -type=Breed

type Breed int

const (
	Poodle Breed = iota
	Beagle
	Chihuahua
	Labrador
	Pug
)

type Dog struct {
	name  string
	breed Breed
}

func (d Dog) String() string {
	return fmt.Sprintf("%s %s", d.name, d.breed)
}

func main() {
	d1 := Dog{
		name:  "Kalu",
		breed: 2,
	}

	fmt.Println(d1)
}
