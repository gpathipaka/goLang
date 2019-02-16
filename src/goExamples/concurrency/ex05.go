package main

import (
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(c chan string) {
	fmt.Println("Service1() start ", time.Since(start))
	time.Sleep(3 * time.Second)
	c <- "Hello From Service 1"
}

func service2(c chan string) {
	fmt.Println("Service2() start ", time.Since(start))
	time.Sleep(3 * time.Second)
	c <- "Hello From Service 2"
}

func main() {
	fmt.Println("main() started", time.Since(start))
	c1 := make(chan string, 2)
	c2 := make(chan string, 2)

	go service1(c1)
	go service2(c2)

	//time.Sleep(3 * time.Second)
	select {
	case res := <-c1:
		fmt.Println("Response from ch1 ", res, time.Since(start))
	case res := <-c2:
		fmt.Println("Response from ch2 ", res, time.Since(start))
	case <-time.After(2 * time.Second):
		fmt.Println("No Response received Timed out.....", time.Since(start))
	default:
		fmt.Println("No Response received", time.Since(start))
	}

	fmt.Println("main() stopeed", time.Since(start))

}
