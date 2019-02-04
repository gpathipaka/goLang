package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("length ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("set:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := c[2:5]
	fmt.Println("l:", l)
	fmt.Println("l:", c[:5])
	fmt.Println("l:", c[2:])
}
