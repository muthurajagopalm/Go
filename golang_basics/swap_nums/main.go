package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("Enter two numbers to swap:")
	fmt.Scan(&a, &b)
	a, b = b, a
	fmt.Println("The swapped numbers are:", a, b)
}
