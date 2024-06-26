package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Starting...")

	result := sumRange(20)

	fmt.Println(result)
}

func sumRange(n int) int {
	work := make(chan int)
	res := make(chan int)

	go func() {
		var sum int
		for {
			select {
			case i, ok := <-work:
				if !ok {
					res <- sum
					return
				}
				sum += i
			}
		}
	}()

	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			work <- i
			wg.Done()
		}()
	}

	wg.Wait()
	close(work)

	return <-res
}
