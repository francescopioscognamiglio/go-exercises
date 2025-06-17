/**
	Create a simple web server using net/http package from
	the standard library. It should serve a plain text web
	page greeting you (e.g., “Hi John Doe!”). The name should
	be sent via query string.
	The URL should be something like
	http://localhost:8090/?name=John+Doe.
**/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		fmt.Fprintf(w, "Hi %s!\n", q.Get("name"))

	}
	http.HandleFunc("/", h1)
	log.Fatal(http.ListenAndServe(":8090", nil))
}

