package main

import (
	"math/rand"
	"fmt"
	"strconv"
)

func passToTheCaller() (string, error) {
	if rand.Intn(2) == 1 {
		_, err := strconv.ParseInt("ciao", 1, 1)
		return "error", fmt.Errorf("error generated %w", err)
	}
	return "no error", nil
}

func main() {
	_, err := passToTheCaller()
	if (err != nil) {
		fmt.Printf("Found the error: %s\n", err.Error())
	} else {
		fmt.Println("No error found")
	}
}

