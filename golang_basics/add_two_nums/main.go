package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("Enter first number:")
	fmt.Scan(&a)
	fmt.Print("Enter second number:")
	fmt.Scan(&b)

	sum := a + b
	fmt.Println("The sum is:", sum)
}
