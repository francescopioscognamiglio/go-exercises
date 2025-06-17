/**
	Create a simple web server using net/http package from
	the standard library. It should serve a plain text web
	page on port 8090 with the output of exercise 2.
	See https://pkg.go.dev/net/http for package documentation.
**/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func isEven(number int) (isEven bool) {
	isEven = number%2 == 0
	return
}

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		randomNumber := rand.Intn(100)
		i := 0
		for ; isEven(randomNumber) == false; i++ {
			fmt.Fprintf(w, "%d - Generated an odd number %d, we have to continue\n", i, randomNumber)
			randomNumber = rand.Intn(100)
		}
		fmt.Fprintf(w, "%d - Generated the first even number %d, the program is stopped\n", i, randomNumber)

	}
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

