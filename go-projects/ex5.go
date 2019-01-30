package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")

	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 101
	fmt.Println("set:", a)
	fmt.Println("get: ", a[4])

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b :", b)

	fmt.Println("length of b ", len(b))

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
