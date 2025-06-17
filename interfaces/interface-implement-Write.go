package main

import (
	"io"
	"math/rand"
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func (v Vertex) Write(content [] byte) (int, error) {
	fmt.Printf("Writing %b", content)
	return 0, nil
}

/**
	Since we defined the method Write on Vertex
	with the same signature on the method io.Writer.Write,
	we can use the type io.Writer instead of the more specific Vertex
**/
func WriteSomething(w io.Writer, content [] byte) {
	w.Write(content)
}

func main() {
	v := Vertex{rand.Intn(10), rand.Intn(10)}
	a := [] byte {0, 1, 0}
	WriteSomething(v, a)
}

