package main

import (
	"encoding/json"
	"fmt"
)

// create something into a json
// a slice of structs

type Person struct {
	First   string
	Last    string
	Details Details
}

type Details struct {
	Phone   int
	Address string
}

func main() {

	//a1 := Person{
	//	First: "John",
	//	Last:  "Smith",
	//	Details: Details{
	//		Phone:   1,
	//		Address: "mars",
	//	},
	//}
	//
	//a2 := Person{
	//	First: "John",
	//	Last:  "watson",
	//	Details: Details{
	//		Phone:   1029,
	//		Address: "ecocity",
	//	},
	//}
	//
	//a3 := Person{
	//	First: "will",
	//	Last:  "Smith",
	//	Details: Details{
	//		Phone:   1123,
	//		Address: "pluto",
	//	},
	//}
	//
	//data := []Person{a1, a2, a3}
	//
	//byteSlice, err := json.Marshal(data)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(string(byteSlice))

	const jsonData = `
[
    {
        "First": "John",
        "Last": "Smith",
        "Details": {
            "Phone": 1,
            "Address": "mars"
        }
    },
    {
        "First": "John",
        "Last": "Watson",
        "Details": {
            "Phone": 1029,
            "Address": "ecocity"
        }
    },
    {
        "First": "Will",
        "Last": "Smith",
        "Details": {
            "Phone": 1123,
            "Address": "pluto"
        }
    }
]
`

	var b2 []Person

	err := json.Unmarshal([]byte(jsonData), &b2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(b2)
}
