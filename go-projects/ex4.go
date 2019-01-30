package main

import "fmt"
import "time"

func main() {
	fmt.Println("Switch Statement in GO")

	i := 2
	fmt.Print("Write ", i, " as")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	fmt.Println("today is ", time.Now().Weekday())

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend")
	default:
		fmt.Println("Its a weekday")
	}

	t := time.Now()
	fmt.Println(t.Hour())

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know the type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("Hey")
}
