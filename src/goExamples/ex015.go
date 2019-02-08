package main

import (
	"fmt"
)

func main() {

	x := 10
	zero(x)
	fmt.Println(x)
	zeroPointer(&x)
	fmt.Println(x)
}

func zero(x int) {
	x = 0
}
func zeroPointer(x *int) {
	*x = 0
}
