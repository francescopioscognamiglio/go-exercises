package main

import (
  "fmt"
)

func main() {
  var channel = make(chan int, 2)
  // set values in a channel via goruotine
  go func() {
    for i := 0; i < 10; i++ {
      channel <- i
      fmt.Printf("capacity of the channel: %d\n", cap(channel))
      fmt.Printf("values in the channel right now: %d\n", len(channel))
    }
  }()

  // read values from channel (and wait if no value has been set)
  for j := 0; j < 10; j++ {
    fmt.Println(<-channel)
    fmt.Printf("capacity of the channel: %d\n", cap(channel))
    fmt.Printf("values in the channel right now: %d\n", len(channel))
  }
}

