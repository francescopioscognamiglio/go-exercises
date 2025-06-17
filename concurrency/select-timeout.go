package main

import (
  "fmt"
  "time"
)

func main() {
  var ch1 = make(chan int)
  var ch2 = make(chan int)
  var ch3 = make(chan int)

  var timeout = time.After(5 * time.Second)

  select {
    case v1 := <- ch1:
      fmt.Printf("Received %d from channel 1\n", v1)
    case v2 := <- ch2:
      fmt.Printf("Received %d from channel 2\n", v2)
    case v3 := <- ch3:
      fmt.Printf("Received %d from channel 3\n", v3)
    // we can set a timeout after which we exited from the select statement
    case <- timeout:
      fmt.Println("No one communicated in 5 seconds")
  }
}

