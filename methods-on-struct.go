package main

import (
	"fmt"
	"math/rand"
)

type Vertex struct {
	X int
	Y int
}

func (v Vertex) Equal(other Vertex) bool {
	return v.X == other.X && v.Y == other.Y
}

func main() {
	v1 := Vertex{rand.Intn(2), rand.Intn(2)}
	v2 := Vertex{rand.Intn(2), rand.Intn(2)}
	b := v1.Equal(v2)
	fmt.Printf("Are {%d,%d} and {%d,%d} equal? %t\n", v1.X, v1.Y, v2.X, v2.Y, b)
}

