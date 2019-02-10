package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sum(x...))
	fmt.Println(evenSum(sum, x...))
	fmt.Println(oddSum(sum, x...))
	fmt.Println("Total Sum = ", evenSum(sum, x...)+oddSum(sum, x...))
}

func sum(x ...int) int {
	sum := 0
	for _, val := range x {
		sum += val
	}
	return sum
}

func sum1(x []int) int {
	sum := 0
	for _, val := range x {
		sum += val
	}
	return sum
}
func evenSum(fn func(x ...int) int, s ...int) int {
	a := []int{}
	for _, val := range s {
		if val%2 == 0 {
			a = append(a, val)
		}
	}
	return fn(a...)

}

func oddSum(fn func(x ...int) int, s ...int) int {
	a := []int{}
	for _, val := range s {
		if val%2 != 0 {
			a = append(a, val)
		}
	}
	return fn(a...)

}
