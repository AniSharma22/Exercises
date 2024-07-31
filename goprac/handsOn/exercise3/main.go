package main

import (
	"fmt"
	"github.com/GoesToEleven/puppy"
)

// $env:GOOS = "darwin"

// go install places your builded executable at GOPATH bin folder

var PackageLevel = 10000 // package level variable
const PackLevel = 20000  // package level constant

func main() {

	a := 10
	fmt.Println(a)

	puppy.Bark()

	fmt.Println("Package level variables: ", PackageLevel, PackLevel)

}
