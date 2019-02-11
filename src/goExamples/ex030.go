package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, playground")
	total := mul3or5(sum, 10)
	fmt.Println("total ", total)
	
	total = mul3or5(sum, 1000)
	fmt.Println("total ", total)
}

func mul3or5(f func(nums ...int) int, n int) int {
	numbers := []int{}
	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			numbers = append(numbers, i)
		}
	}
	fmt.Println(numbers)
	return f(numbers...)
}

func sum(nums ...int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}
