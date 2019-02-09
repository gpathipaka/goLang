package main

import "fmt"

type Person struct {
	firstName, lastName string
	favIceCreamFlavs    []string
}

func main() {
	fmt.Println("Hello")
	m := make(map[string]Person)
	p1 := Person{
		firstName:        "Gangadhar",
		lastName:         "Pathipaka",
		favIceCreamFlavs: []string{"buttorScotch", "Vennila"},
	}

	m[p1.lastName] = p1

	for k, v := range m {
		fmt.Println(k)
		for i, val := range v.favIceCreamFlavs {
			fmt.Printf("\t %v \t %v\n", i, val)
		}
	}
}
