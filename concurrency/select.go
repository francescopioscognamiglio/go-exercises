package main

import (
  "fmt"
)

func main() {
  var ch1 = make(chan int)
  var ch2 = make(chan int)
  var ch3 = make(chan int)

  go func() {
    ch2 <- 5
  }()

  select {
    case v1 := <- ch1:
      fmt.Printf("Received %d from channel 1\n", v1)
    case v2 := <- ch2:
      fmt.Printf("Received %d from channel 2\n", v2)
    case v3 := <- ch3:
      fmt.Printf("Received %d from channel 3\n", v3)
  }
}

