package main

import "fmt"

var PackageLevel = 10000 // packge level variable
const PackLevel = 20000  // package level constant

func main() {

	const (
		_  = iota
		kb = 1 << (10 * iota)
		mb = 1 << (10 * iota)
		gb = 1 << (10 * iota)
		tb = 1 << (10 * iota)
	)

	var a int // zero value = 0
	b := 20

	var x, y, z int = 1, 2, 3 //multiple initialisations

	_ = a
	_ = b

	var a1 string = "ash"
	var a2 int = 10
	var a3 float64 = 3.44

	fmt.Printf("%T\t%v", a1, a2, a3)

	var num1, num2, num3 int = 747, 911, 90210

	fmt.Printf("%d \t %b \t %x", num1)
	fmt.Printf("%d \t %b \t %x", num2)
	fmt.Printf("%d \t %b \t %x", num3)

	var largestInt8 int8 = 127
	var largestUint8 uint8 = 255

	fmt.Println(largestUint8, largestInt8)

	fmt.Println("Package level variables: ", PackageLevel, PackLevel)

}
