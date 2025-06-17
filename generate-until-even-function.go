/**
	Define a function that, given an integer,
	returns true if the number is even, false otherwise.
**/
package main

import (
	"fmt"
	"math/rand"
)

func isEven(number int) (isEven bool) {
	isEven = number%2 == 0
	return
}

func main() {
	randomNumber := rand.Intn(100)
	for isEven(randomNumber) == false {
		fmt.Printf("Generated an odd number %d, we have to continue\n", randomNumber)
		randomNumber = rand.Intn(100)
	}
	fmt.Printf("Generated the first even number %d, the program is stopped\n", randomNumber)
}

