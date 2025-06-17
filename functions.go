package main

import "fmt"

func sumsub(n1 int, n2 int) (int, int) {
	return n1 + n2, n1 - n2
}

func sumsubadv(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func fact(n int) int {
	if n <= 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func main() {
	sum, sub := sumsub(3, 2)
	fmt.Println(sum)
	fmt.Println(sub)

	sum, sub = sumsubadv(3, 2)
	fmt.Println(sum)
	fmt.Println(sub)

	fmt.Println(fact(10))
}

