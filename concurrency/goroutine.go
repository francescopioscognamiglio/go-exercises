package main

import (
  "fmt"
  "time"
)

func printNumbers(start int, end int, asc bool) {
  if asc {
    for i := start; i < end; i++ {
      fmt.Println(i)
    }
  } else {
    for i := end; i > start; i-- {
      fmt.Println(i)
    }
  }
}

func main() {
  // print numbers from 0 to 9
  go printNumbers(0, 10, true)
  // print numbers from 10 to 1
  printNumbers(0, 10, false)
  time.Sleep(100 * time.Millisecond)
}

