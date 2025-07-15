package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		ch1 <- "You have a new message!"
	}()

	select {
	case msg := <-ch1:
		fmt.Println("You have received a message:", msg)
	case <-time.After(4 * time.Second):
		fmt.Println("No message received within the timeout period.")
	}
}
