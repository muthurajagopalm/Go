package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	ch := make(chan int)

	go func() {
		val := <-ch
		fmt.Println("Value received from channel:", val)
		wg.Done()

	}()

	go func() {
		val := <-ch
		fmt.Println("Value received from channel:", val)
		wg.Done()

	}()

	go func() {
		val := <-ch
		fmt.Println("Value received from channel:", val)
		wg.Done()
	}()
	ch <- 27
	ch <- 36
	ch <- 45
	wg.Wait()
}
