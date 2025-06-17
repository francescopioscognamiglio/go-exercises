/**
	Write a Go program that generates random numbers
	until you find an even number.
**/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(100)
	for randomNumber%2 == 1 {
		fmt.Printf("Generated an odd number %d, we have to continue\n", randomNumber)
		randomNumber = rand.Intn(100)
	}
	fmt.Printf("Generated the first even number %d, the program is stopped\n", randomNumber)
}

