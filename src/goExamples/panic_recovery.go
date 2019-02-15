package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	wg.Add(1)
	fmt.Println("Main() start")
	go foo(c)

	for v := range c {
		fmt.Println(v)
	}
	wg.Wait()
	fmt.Println("Main() end")
}

func foo(c chan int) {
	defer wg.Done()

	//Defer 3
	defer func() {
		fmt.Println("test defer....")
	}()

	//Defer 2
	defer func() {
		fmt.Println("closing connection")
		close(c)
	}()

	//Defer 1
	defer func() {
		if err := recover(); err != nil {
			c <- 10
			fmt.Println("Recovered this ", err)
		}
		fmt.Println("Exiting the recover block")
	}()
	c <- 5
	panic("foo: failed......")
}
