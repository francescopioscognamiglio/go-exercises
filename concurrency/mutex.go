package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	countwords map[string]int
}

// Increment increments the counter for the given key.
func (c *SafeCounter) Increment(key string) {
	// Lock so only one goroutine at a time can access the map
  c.mu.Lock()
	c.countwords[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	// Lock so only one goroutine at a time can access the map
  c.mu.Lock()
	defer c.mu.Unlock()
	return c.countwords[key]
}

func main() {
	c := SafeCounter{countwords: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Increment("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
