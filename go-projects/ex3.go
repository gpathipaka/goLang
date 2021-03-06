package main

import (
	"fmt"
)

func main() {

	forLoop()

	ifElse()

}

func ifElse() {
	fmt.Println("************************")
	if 7%2 == 0 {
		fmt.Println("7 is a even number")
	} else {
		fmt.Println("7 is a odd number")
	}

	if 8%2 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative number")
	} else if num < 10 {
		fmt.Println(num, " has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
	fmt.Println("************************")
}

func forLoop() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
