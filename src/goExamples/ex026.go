package main

import "fmt"

type person struct {
	first, last string
	age         int
}

func (p person) speak() {
	fmt.Println("My Name is", p.first, p.last, "And My age is", p.age)
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
		age:   28,
	}
	p.speak()
}
