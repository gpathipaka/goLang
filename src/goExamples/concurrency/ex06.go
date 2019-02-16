package main

import (
	"fmt"
	"runtime"
)

func puts(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)

}
func main() {

	c := make(chan int)

	go puts(c)

	for v := range c {

		fmt.Println(v)
	}

	/*
			for v := range c {
				fmt.Println(v)
			}

		for {

			select {
			case res := <-c:
				fmt.Println("Received from channel n ", res)
			default:
				fmt.Println("No Response from Chanel ")
			}
		}

	*/

	fmt.Println(runtime.NumGoroutine())
}
