package main

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
)

// niladic means no arguments in reference to a function
// generates a random number from 0 to 249

func Intn(n int) int {
	return rand.Intn(n)
}

func Print99(flag bool) {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	if flag {
		fmt.Println(100)
	}
}

func Intn2(n int) (int, int) {
	return rand.Intn(n), rand.Intn(n)
}

func init() {
	fmt.Println("this runs before main function")
}

func init() {

	fmt.Println("this is init2")
}
func main() {
	// Define the PowerShell command to run Get-FileHash
	psCommand := `Get-FileHash -Algorithm SHA256 -Path "./data.txt" | Format-List`

	// Create the PowerShell command
	cmd := exec.Command("powershell", "-Command", psCommand)

	// Run the command and capture the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}

	// Print the output
	fmt.Printf("SHA-256 Checksum:\n%s", output)

	x := Intn(250)
	fmt.Println("value of x is", x)

	if x > 0 && x <= 100 {
		fmt.Println("value of x is between 0 and 100")
	} else if x > 100 && x <= 200 {
		fmt.Println("value of x is between 100 and 200")
	} else {
		fmt.Println("value of x is between 200 and 250")
	}

	// similar thing can be achieved using switch

	switch {
	case x > 0 && x <= 100:
		fmt.Println("value of x is between 0 and 100")
	case x > 100 && x <= 200:
		fmt.Println("value of x is between 100 and 200")
	case x > 200 && x <= 250:
		fmt.Println("value of x is between 250 and 250")
	default:
		fmt.Println("unexpected behaviour")

	}

	a1, a2 := Intn2(10)
	fmt.Println(a1, a2)

	if x < 4 && y < 4 {
		fmt.Println("value of x and y both is less than4")
	} else if x > 6 && y > 6 {
		fmt.Print("value of x and y both is greater than 6")
	} else if x >= 4 && y <= 6 {
		fmt.Print("value of x and y both ")
	} else if y != 5 {
		fmt.Print("y is not 5")
	} else {
		fmt.Println("no cases are met")
	}

	var a23 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(a23)
}
