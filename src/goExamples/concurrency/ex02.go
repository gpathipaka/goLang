package main

import "fmt"

func main() {
	c := make(chan int)

	//send
	go foo(c)

	//receive
	bar(c)

	fmt.Println("done....")
}

func foo(c chan<- int) {
	c <- 55
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
