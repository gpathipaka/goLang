package main

import "fmt"

func main() {

	s := struct {
		lastName  string
		friends   map[string]int
		favDrinks []string
	}{
		lastName: "Ishu",
		friends: map[string]int{
			"A": 123,
			"B": 444,
		},
		favDrinks: []string{"Martini", "beer"},
	}

	fmt.Println(s)
}
