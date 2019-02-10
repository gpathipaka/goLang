package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Println("x before foo(x) ", x)
	foo(x)
	fmt.Println("x after foo(x) ", x)

	fmt.Println("x before bar(&x) ", x)
	bar(&x)
	fmt.Println("x after bar(&x) ", x)
}

func foo(x int) {
	x = 100
	fmt.Println("x inside foo(x int) ", x)
}
