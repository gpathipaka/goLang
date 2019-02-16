package main

import (
	"fmt"
	"strings"
)

func main() {
	//ch := make(chan int, 1)
	ch := make(chan int)
	go func() {
		ch <- 100
		ch <- 101
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(strings.Repeat("*", 50))
	fmt.Printf("Type of Ch : %T\n", ch)
}
