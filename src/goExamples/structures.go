package main

import (
	"fmt"
)

type student struct {
	id                  int
	lastName, firstName string
}

func main() {
	var ishan student
	ishan.id = 001
	ishan.lastName = "Ishan"
	ishan.firstName = "Pathipaka"
	fmt.Println(ishan.id)
	fmt.Println(ishan.lastName)
	fmt.Println(ishan.firstName)

	reyan := student{
		firstName: "Reyan",
		lastName:  "Pathipaka",
		id:        2,
	}
	fmt.Println(reyan)
	
	ganga := student {3, "Gangadhar", "Pathipaka"}
	fmt.Println(ganga)
}
