package main

import (
	"fmt"
)

// creating my own error type MyError
type MyError struct {
  LineNumber uint
  Message string
}

// implementing the function Error of the built-in interface error
func (e *MyError) Error() string {
  return fmt.Sprintf("Error in line %d: %s", e.LineNumber, e.Message)
}

// returning the built-in interface error using my own error type MyError
func checkLine() error {
  return &MyError{LineNumber: 15, Message: "Missing semicolon"}
}

func main() {
  err := checkLine()
  if err != nil {
    fmt.Println(err)
  }
}

