package main

import "fmt"

type Person struct {
	firstName, lastName string
	favIceCreamFlavs    []string
}

func main() {
	fmt.Println("Hello")

	p1 := Person{
		firstName:        "Gangadhar",
		lastName:         "Pathipaka",
		favIceCreamFlavs: []string{"buttorScotch", "Vennila"},
	}

	for i, v := range p1.favIceCreamFlavs {
		fmt.Println(i, v)
	}
	fmt.Println(p1)

}
