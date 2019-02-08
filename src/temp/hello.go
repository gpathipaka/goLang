package main

import "fmt"

func main() {
	fmt.Println("Start here ")
	foo()
	fmt.Println("almost done...")
	bar()
}

func foo() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func bar() {
	fmt.Println("done....")
}
