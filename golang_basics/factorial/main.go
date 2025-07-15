package main

import (
	"fmt"
)

func main() {
	var n, fact int = 5, 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	fmt.Println(fact)

}
