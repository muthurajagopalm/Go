package main

import (
	"fmt"
	"math/rand"
)

func main() {
	quotes := []string{
		"Believe in yourself!",
		"Live and let live",
		"Hard work beats talent",
		"Keep pushing forward",
		"Love yourself!",
	}

	random_quotes := rand.Intn(len(quotes))
	fmt.Println("Quotes of the moment")
	fmt.Println(quotes[random_quotes])

}
