package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Print("Enter a number: ")
	fmt.Scan(&num)
	if num%2 == 0 {
		fmt.Printf("%d is an even number. \n", num)
	} else {
		fmt.Printf("%d is an odd number. \n", num)
	}
}
