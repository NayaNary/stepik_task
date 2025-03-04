// Конкурентно-безопасный стек.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// начало решения

// Stack представляет конкурентно-безопасный стек без блокировок.
type Stack struct {
	top atomic.Pointer[Node]
}

// Node представляет элемент стека.
type Node struct {
	val  int
	past *Node
}

// Push добавляет значение на вершину стека.
func (s *Stack) Push(val int) {
	for {
		current := s.top.Load()
		node := &Node{val: val, past: current}
		res := s.top.CompareAndSwap(current, node)
		if res {
			break
		}
	}
}

// Pop удаляет и возвращает вершину стека.
// Если стек пуст, возвращает false.
func (s *Stack) Pop() (int, bool) {
	for {
		current := s.top.Load()
		if current == nil {
			return 0, false
		}
		res := s.top.CompareAndSwap(current, current.past)
		if res {
			return current.val, res
		}
	}
}

// конец решения

func main() {
	var wg sync.WaitGroup
	wg.Add(1000)

	stack := &Stack{}
	for i := range 1000 {
		go func() {
			time.Sleep(time.Millisecond)
			stack.Push(i)
			wg.Done()
		}()
	}

	wg.Wait()

	// val := stack.top.Load()
	// for range 1000 {
	// 	if val != nil {
	// 		fmt.Println("for", val.val)
	// 	}
	// 	if val.past == nil {
	// 		break
	// 	}
	// 	val = val.past
	// }

	count := 0
	for _, ok := stack.Pop(); ok; _, ok = stack.Pop() {
		count++
	}
	fmt.Println(count)
}
