/**
	Write a Go program that prints
	if the current "second" is even or odd.
**/
package main

import (
	"fmt"
	"time"
)

func main() {
	currentSecond := time.Now().Second()
	if currentSecond%2 == 0 {
		fmt.Printf("The current second %d is even\n", currentSecond)
	} else {
		fmt.Printf("The current second %d is odd\n", currentSecond)
	}
}

