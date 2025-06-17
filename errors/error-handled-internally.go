package main

import (
	"math/rand"
	"fmt"
)

func handleErrorInternally() (string) {
	if rand.Intn(2) == 0 {
		return "no error"
	} else {
		return "error"
	}
}

func main() {
	s := handleErrorInternally()
	fmt.Printf("Found any error? %s\n", s)
}

