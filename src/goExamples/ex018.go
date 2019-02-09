package main

import "fmt"

type vehicle struct {
	doors, color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t := truck{
		vehicle: vehicle{
			doors: "4",
			color: "silver",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: "4",
			color: "white",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(s)
}
