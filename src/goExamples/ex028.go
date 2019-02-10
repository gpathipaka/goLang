package main

import "fmt"

func main() {

	func() {
		fmt.Println("Ananymous function")
	}()

	fn := func() int {
		return 999
	}
	fmt.Println(fn())

	x := incrementByOne()
	fmt.Printf("Type of x is %T\n", x)
	fmt.Println("X", x())
	fmt.Println("X ", x())
	fmt.Println("X ", x())
	y := incrementByOne()
	fmt.Println("Y ", y())
	fmt.Println("Y ", y())
	fmt.Println("X ", x())
	fmt.Println("X ", x())
	fmt.Println("Y ", y())
	fmt.Println("Y ", y())

}

func incrementByOne() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
