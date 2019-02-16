package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		go f(i)
	}

	var input string
	fmt.Scanln("Please enter something....", &input)
	fmt.Println(input)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		time.Duration(rand.Intn(250))
	}
}
