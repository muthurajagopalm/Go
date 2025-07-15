package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Hello from goroutine %d\n", n)
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines completed.")
}
