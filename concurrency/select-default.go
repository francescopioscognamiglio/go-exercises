package main

import (
  "fmt"
)

func main() {
  var ch1 = make(chan int)
  var ch2 = make(chan int)
  var ch3 = make(chan int)

  select {
    case v1 := <- ch1:
      fmt.Printf("Received %d from channel 1\n", v1)
    case v2 := <- ch2:
      fmt.Printf("Received %d from channel 2\n", v2)
    case v3 := <- ch3:
      fmt.Printf("Received %d from channel 3\n", v3)
    // we can also set a default in case all the channels are empty
    default:
      fmt.Println("No one is ready to communicate")
  }
}

