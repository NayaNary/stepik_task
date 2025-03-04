// Блокирующая очередь.
package main

import (
	"fmt"
	"sync"
)

// начало решения

// Queue - блокирующая FIFO-очередь.
type Queue struct {
	// TODO: переделать на срез и sync.Cond.
	items []int
	cond  sync.Cond
}

// NewQueue создает новую очередь.
func NewQueue() *Queue {
	// TODO: очередь должна быть безразмерной.
	return &Queue{
		items: make([]int, 0),
		cond:  *sync.NewCond(&sync.Mutex{}),
	}
}

// Put добавляет элемент в очередь.
// Поскольку очередь безразмерная, никогда не блокируется.
func (q *Queue) Put(item int) {
	q.cond.L.Lock()
	var res []int
	res = append(res, item)
	res = append(res, q.items...)
	// q.items = slices.Insert(q.items, 0, item)
	q.items = res
	q.cond.Signal()
	q.cond.L.Unlock()
}

// Get извлекает элемент из очереди.
// Если очередь пуста, блокируется до момента,
// пока в очереди не появится элемент.
func (q *Queue) Get() int {
	q.cond.L.Lock()
	for len(q.items) == 0 {
		q.cond.Wait()
	}
	res := q.items[len(q.items)-1]
	q.items = q.items[:len(q.items)-1]
	q.cond.L.Unlock()
	return res
}

// Len возвращает количество элементов в очереди.
func (q *Queue) Len() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	return len(q.items)
}

// конец решения

func main() {
	var wg sync.WaitGroup
	q := NewQueue()

	wg.Add(1)
	go func() {
		for i := range 100 {
			q.Put(i)
		}
		wg.Done()
	}()
	wg.Wait()

	total := 0

	wg.Add(1)
	go func() {
		for range 100 {
			total += q.Get()
		}
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("Put x100, Get x100, Total:", total)
}
