package main

import "fmt"

type person struct {
	lastName, firstName string
}

type secretAgent struct {
	person
	ltk bool
}

func main() {

	sa1 := secretAgent{
		person: person{
			lastName:  "Pathipaka",
			firstName: "Ishan",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			lastName:  "Pathipaka",
			firstName: "Reyan",
		},
		ltk: false,
	}
	sa1.speak()
	sa2.speak()
}

func (s secretAgent) speak() {
	fmt.Println(s)
}
