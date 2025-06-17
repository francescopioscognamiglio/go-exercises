package main

import "fmt"

func main() {
	defer fmt.Println(", boy!")
	defer fmt.Print(" world")
	fmt.Print("hello")
}

