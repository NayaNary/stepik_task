// Безопасная карта
package main

import (
	"fmt"
	"sync"
)

// начало решения

// Counter представляет безопасную карту частот слов.
// Ключ - строка, значение - целое число.
type Counter struct {
	mx         sync.Mutex
	mapCountre map[string]int
}

// Increment увеличивает значение по ключу на 1.
func (c *Counter) Increment(str string) {
	c.mx.Lock()
	c.mapCountre[str] += 1
	c.mx.Unlock()
}

// Value возвращает значение по ключу.
func (c *Counter) Value(str string) int {
	c.mx.Lock()
	defer c.mx.Unlock()
	return c.mapCountre[str]
}

// Range проходит по всем записям карты,
// и для каждой вызывает функцию fn, передавая в нее ключ и значение.
func (c *Counter) Range(fn func(key string, val int)) {
	c.mx.Lock()
	defer c.mx.Unlock()
	for key, val := range c.mapCountre {
		fn(key, val)
	}
}

// NewCounter создает новую карту частот.
func NewCounter() *Counter {
	return &Counter{
		mapCountre: make(map[string]int),
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

	fmt.Println("two:", counter.Value("two"))
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Print("{ ")
		counter.Range(func(key string, val int) {
			fmt.Printf("%s:%d ", key, val)
		})
		fmt.Println("}")
	}()
	go func() {
		defer wg.Done()
		counter.Increment("two")
	}()

	wg.Wait()

	fmt.Print("{ ")
	counter.Range(func(key string, val int) {
		fmt.Printf("%s:%d ", key, val)
	})
	fmt.Println("}")
}
