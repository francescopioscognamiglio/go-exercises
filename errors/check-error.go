package main

import (
  "fmt"
  "errors"
)

var ErrorTooLarge = errors.New("the value is too large")

func CheckSmallNumber(n int) (int, error) {
  if n > 100 {
    return 0, ErrorTooLarge
  }
  return n, nil
}

func main() {
  _, err := CheckSmallNumber(126)
  if errors.Is(err, ErrorTooLarge) {
    fmt.Println("The given number is too large")
  }
}

