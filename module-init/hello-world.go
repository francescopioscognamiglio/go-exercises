/**
	Hello World in a module.
**/
package main

import (
	"fmt"
	"utils/writer"
)

func main() {
	v := "Pio"
	writer.SayHelloWorld(v)
	writer.SayGoodBye(v)
}

