package main

import (
	"fmt"
)

func main() {

	/*
		panic("YES PANIC")
		str := recover()
		fmt.Println(str)

	*/

	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("PANIC")

}
