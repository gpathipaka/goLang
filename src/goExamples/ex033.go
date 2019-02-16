package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs ", runtime.NumCPU())
	fmt.Println("Num Goroutines ", runtime.NumGoroutine())
	const gs = 100
	counter := 0
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i <= 100; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()

		}()
		fmt.Println("Num Goroutines ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Num Goroutines ", runtime.NumGoroutine())
	fmt.Println("Counter is ", counter)
}
