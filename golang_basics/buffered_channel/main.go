package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 2)

	ch <- 27
	ch <- 36

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
