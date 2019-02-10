package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Square struct {
	side float64
}

type shape interface {
	area() float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s Square) area() float64 {
	return s.side * s.side
}

func info(s shape) {
	switch s.(type) {
	case Circle:
		fmt.Println("Area of a circle with r ", s.(Circle).radius, "is", s.area())
	case Square:
		fmt.Println("Area of a Square with r ", s.(Square).side, "is", s.area())
	default:
		fmt.Println("you are in wrong place")

	}
}

func main() {
	c := Circle{5}
	s := Square{5}

	info(c)
	info(s)
}
