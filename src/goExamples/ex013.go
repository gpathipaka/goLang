package main

import "fmt"

func main() {
	fmt.Println("start")
	defer first()
	second()
	fmt.Println("done")
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}
