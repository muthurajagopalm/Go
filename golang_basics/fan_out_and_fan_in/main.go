package main

import (
	"fmt"
)

func worker(input <-chan int, output chan<- int) {
	for num := range input {
		squares := num * num
		output <- squares
	}
}

func main() {
	input := make(chan int)
	output := make(chan int)

	for i := 0; i < 3; i++ {
		go worker(input, output)
	}

	go func() {
		for i := 1; i <= 5; i++ {
			input <- i
		}
		close(input)
	}()

	for i := 0; i < 5; i++ {
		result := <-output
		fmt.Println("Square:", result)
	}
}
