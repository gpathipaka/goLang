package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First, Last string
	Age         int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   35,
	}

	p2 := person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   28,
	}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println("error : ", err)
	}

	str := string(bs)
	fmt.Println(str)

	//unmarshal

	var p []person
	err = json.Unmarshal(bs, &p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(p)

}
