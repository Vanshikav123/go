package main

import (
	"fmt"
	"sync"
	"time"
)

//safeCounter is safe to use concurrently
/*A Mutex is a method used as a locking mechanism to ensure that only one Goroutine is accessing the critical section of code at any point of time.
This is done to prevent race conditions from happening. Sync package contains the Mutex. Two methods defined on Mutex

Lock
Unlock
Any code present between a call to Lock and Unlock will be executed by only one Goroutine.

mutex.Lock()

x = x + 1 // this statement be executed
          // by only one Goroutine
          // at any point of time

mutex.Unlock()
If one Goroutine already has the lock and if a new Goroutine is trying to get the lock, then the new Goroutine will be stopped until the mutex is unlocked. To understand this concept, let’s first understand a program having race conditions.*/

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increents the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	//Lock so only one goroutine at a time can access the mutex while holding the counter lock
	c.v[key]++
	c.mu.Unlock()
}

// value returns the curent value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i <= 500; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
