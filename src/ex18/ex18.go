package main

import (
	"fmt"
	"sync"
)

type counter struct {
	value int
	mu    sync.Mutex
}

func (c *counter) increment() {
	c.mu.Lock()
	c.value++
	defer c.mu.Unlock()
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	count := &counter{}
	amountGoroutine := 30
	var wg sync.WaitGroup

	wg.Add(amountGoroutine)
	for i := 0; i < amountGoroutine; i++ {
		go func() {
			count.increment()
			defer wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("Result value:", count.value)
}
