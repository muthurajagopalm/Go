package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Printf("The maximum number is: %d\n", a)
	} else if b > a {
		fmt.Printf("The maximum number is: %d\n", b)
	} else {
		fmt.Println("Both numbers are equal.")
	}

}
