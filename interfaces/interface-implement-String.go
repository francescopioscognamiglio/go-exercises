package main

import (
	"fmt"
	"math/rand"
)

type Vertex struct {
	X int
	Y int
}

func (v Vertex) String() string {
	return fmt.Sprintf("(%d,%d)", v.X, v.Y) // return a string
}

/**
	Since we defined the method String on Vertex
	with the same signature on the method fmt.Stringer.String,
	we can use the type fmt.Stringer instead of the more specific Vertex
**/
func printSomething(s fmt.Stringer) {
	fmt.Println(s.String())
}

func main() {
	v := Vertex{rand.Intn(10), rand.Intn(10)}
	printSomething(v)
}

