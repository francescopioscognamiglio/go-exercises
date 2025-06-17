package main

import "fmt"

func main() {
	// for statement
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	// for like a while
	i := 0
	for i < 10 {
		fmt.Print(i)
		i++
	}
	fmt.Println()

	// if statement
	n := 50
	if n > 10 {
		fmt.Println("n is greater than 10")
	} else if n > 5 {
		fmt.Println("n is greater than 5")
	} else {
		fmt.Println("n is less than 5")
	}

	// switch statement
	switch n {
	case 10:
		fmt.Println("n is equal to 10")
	case 20:
		fmt.Println("n is equal to 20")
	case 30:
		fmt.Println("n is equal to 30")
	default:
		fmt.Println("n is not 10, 20 or 30")
	}

	// switch like an if
	switch {
	case n > 10:
		fmt.Println("n is greater than 10")
	case n > 5:
		fmt.Println("n is greater than 5")
	default:
		fmt.Println("n is less than 5")
	}

}

