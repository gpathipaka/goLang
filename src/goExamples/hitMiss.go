package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var wg sync.WaitGroup

func main() {
	fmt.Println("main() started....")
	wg.Add(2)
	court := make(chan int)

	go player("Ishan", court)
	go player("Reyan", court)

	court <- 1
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()
	for {
		ball, ok := <-court
		// if channel closed....
		if ok == false {
			fmt.Printf("Player %s Won \n", name)
			return
		}

		n := rand.Intn(100)
		if n == 56 {
			fmt.Printf("Player %s Missed \n", name)
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++
		court <- ball
	}

}
