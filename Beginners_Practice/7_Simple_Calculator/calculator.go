package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//check the number of arguments
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run calculator.go <num1> <+|-|*|/> <num2>")
		return
	}
	//convert the number from text to integer
	a, err1 := strconv.Atoi(os.Args[1])
	b, err2 := strconv.Atoi(os.Args[3])
	if err1 != nil || err2 != nil {
		fmt.Println("Enter the valid input")
		return
	}
	//read the operator

	op := os.Args[2]

	//perform the calculation
	switch op {
	case "+":
		fmt.Println("Result:", a+b)
	case "-":
		fmt.Println("Result:", a-b)
	case "*":
		fmt.Println("Result:", a*b)
	case "/":
		fmt.Println("Result:", a/b)
	default:
		fmt.Println("Unknown operator")
	}

}
