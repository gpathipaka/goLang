package main

import "fmt"

type person struct {
	lastName, firstName string
}

type secretAgent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println(s)
}

func (p person) speak() {
	fmt.Println(p)
}

func foo(h human) {
	switch h.(type) {
	case person:
		fmt.Println("you are a person ", h.(person).lastName)
	case secretAgent:
		fmt.Println("you are a secret Agent ", h.(secretAgent).lastName, " And you have ltk ", h.(secretAgent).ltk)
		sa := h.(secretAgent)
		fmt.Println("you are a secret Agent ", sa.firstName, " And you have ltk ", h.(secretAgent).ltk)
	default:
		fmt.Println("you are in wrong place....")
	}
	fmt.Println("Speaking ", h)
}

func main() {

	p1 := person{
		lastName:  "Pathipaka",
		firstName: "Ishan",
	}

	p2 := person{
		lastName:  "Pathipaka",
		firstName: "Reyan",
	}

	p3 := person{
		lastName:  "Romo",
		firstName: "Tony",
	}
	sa1 := secretAgent{
		person: p1,
		ltk:    true,
	}

	sa2 := secretAgent{
		person: p3,
		ltk:    false,
	}
	sa1.speak()
	sa2.speak()
	p2.speak()

	foo(p1)
	foo(sa2)
}
