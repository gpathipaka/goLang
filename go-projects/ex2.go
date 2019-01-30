package main

import (
	"fmt"
	"math"
)

const s string = "this is a constant string"

func main() {
	test()
}
func test() {
	fmt.Println(s)

	const n = 50000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
