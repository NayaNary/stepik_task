package main

import (
	"fmt"
	"sync"
)

// начало решения

type Counter struct {
	*sync.Mutex
	mapRange map[string]int
}

func (c *Counter) Increment(str string) {
	c.Lock()
	defer c.Unlock()
	c.mapRange[str]++
}

func (c *Counter) Value(str string) int {
	c.Lock()
	defer c.Unlock()
	return c.mapRange[str]
}

func (c *Counter) Range(fn func(key string, val int)) {
	c.Lock()
	defer c.Unlock()
	for key, val := range c.mapRange {
		fn(key, val)
	}
}

func NewCounter() *Counter {
	return &Counter{
		mapRange: make(map[string]int),
		Mutex: &sync.Mutex{},
	}
}

// конец решения

func main() {
	counter := NewCounter()

	var wg sync.WaitGroup
	wg.Add(3)

	increment := func(key string, val int) {
		defer wg.Done()
		for ; val > 0; val-- {
			counter.Increment(key)
		}
	}

	go increment("one", 100)
	go increment("two", 200)
	go increment("three", 300)

	wg.Wait()

	go fmt.Println("two:", counter.Value("two"))

	fmt.Print("{ ")
	counter.Range(func(key string, val int) {
		fmt.Printf("%s:%d ", key, val)
	})
	fmt.Println("}")
}
