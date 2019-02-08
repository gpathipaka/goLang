package main

import (
	"fmt"
)

var a int
var b string
var c bool

func main() {

	ex3()
}

func ex3() {
	type hotdog int
	var x hotdog
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}

func ex2() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func ex1() {

	x := 42
	y := "James Bond"
	z := true
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println("Values from Sprint", s)
}
