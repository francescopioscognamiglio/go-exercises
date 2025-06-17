package main

import "fmt"

var x, y = 1.5, 3.5

func main() {
	// this is a comment
	/* this is another comment
	on multiple lines*/

	var a, i = "ciao", 12
	fmt.Println(x, y)
	fmt.Println(a, i)

	// declaration + initialization
	di1 := true
	di2 := "pippo"
	di3, di4 := 1.5, 3.5
	fmt.Println(di1, di2, di3, di4)

	var v [5]int
	v[1] = 3
	v[4] = 10
	fmt.Println(v)

	var t = v[1:3]
	fmt.Println(t)

	var va = v[0:5]
	var vb = v[0:]
	var vc = v[:5]
	var vd = v[:]
	fmt.Println(v)
	fmt.Println(va, vb, vc, vd)

	var x1 = 3
	var x2 = x1
	x2 = 10
	fmt.Println(x1, x2)

	v[0] = 100
	fmt.Println(v)
	fmt.Println(va, vb, vc, vd)

	va[2] = 200
	fmt.Println(v)
	fmt.Println(va, vb, vc, vd)

	var temperatures map[string]int = make(map[string]int)
	fmt.Println(temperatures)
	temperatures["Rome"] = 27
	temperatures["Milan"] = 23
	temperatures["San Francisco"] = 20
	fmt.Println(temperatures)

	var temp = map[string]int{
		"Rome":  27,
		"Milan": 20,
	}
	fmt.Println(temp)

	elem, ok := temperatures["London"]
	fmt.Println(elem, ok)


	type Vertex struct {
		x int
		y int
	}
	structv := Vertex{1, 2}
	fmt.Println(structv)   // print struct
	fmt.Println(structv.x) // print x of the struct
	fmt.Println(structv.y) // print y of the struct
}

