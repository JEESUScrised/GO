package main

import (
	"fmt"
	"sync"
	"time"
)

type call struct {
	done   chan struct{}
	result int
}

type Cache struct {
	mu       sync.Mutex
	values   map[int]int
	inFlight map[int]*call
}

func NewCache() *Cache {
	return &Cache{
		values:   make(map[int]int),
		inFlight: make(map[int]*call),
	}
}

func (c *Cache) Get(key int) int {
	c.mu.Lock()

	if value, ok := c.values[key]; ok {
		c.mu.Unlock()
		return value
	}

	if running, ok := c.inFlight[key]; ok {
		c.mu.Unlock()
		<-running.done
		return running.result
	}

	current := &call{done: make(chan struct{})}
	c.inFlight[key] = current
	c.mu.Unlock()

	result := expensiveFib(key)

	c.mu.Lock()
	current.result = result
	c.values[key] = result
	delete(c.inFlight, key)
	close(current.done)
	c.mu.Unlock()

	return result
}

func expensiveFib(n int) int {
	time.Sleep(1 * time.Second)
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	cache := NewCache()

	go func() { fmt.Println(cache.Get(10)) }()
	go func() { fmt.Println(cache.Get(10)) }()

	time.Sleep(2 * time.Second)
}
