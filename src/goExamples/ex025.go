package main

import (
	"fmt"
)

type person struct {
	lastName, firstName string
}

func changeMe(p *person) {
	(*p).firstName = "Ishan"
	fmt.Println("P inside changeMe ", *p)
}
func main() {
	p := person{"Pathipaka", "Gangadhar"}
	fmt.Println("P befor changeMe ", p)
	changeMe(&p)
	fmt.Println("P after changeMe ", p)

}
