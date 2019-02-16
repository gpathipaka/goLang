package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)

}

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "PING"
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "PONG"
	}
}

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
