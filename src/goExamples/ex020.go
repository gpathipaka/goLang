package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3, 10)
	s = []string{"A", "B", "C", "D", "E"}
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s = append(s, "F", "G", "H", "I")
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s = append(s, "K", "L", "M", "N")
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	for i, v := range s {
		fmt.Println(i, v)
	}
}
