package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	numGoroutines := 3
	chunkSize := len(arr) / numGoroutines
	sumChannel := make(chan int, numGoroutines)
	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == numGoroutines-1 {
			end = len(arr) // Last goroutine takes the rest of the array
		}
		wg.Add(1)

		go func(subArr []int) {
			defer wg.Done()
			partialSum := 0
			for _, val := range subArr {
				partialSum += val
				//fmt.Println("Partial Sum:", partialSum)

			}

			sumChannel <- partialSum

		}(arr[start:end])
	}
	go func() {
		wg.Wait()
		close(sumChannel)
	}()

	totalSum := 0
	for part := range sumChannel {
		totalSum += part

	}

	fmt.Println("Total Sum:", totalSum)

}

// Note: The WaitGroup and closing of the channel should be outside the loop
