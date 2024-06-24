package main

import (
	"errors"
	"fmt"
)

var ErrFull = errors.New("Queue is full")
var ErrEmpty = errors.New("Queue is empty")

// начало решения

// Queue - FIFO-очередь на n элементов
type Queue chan int

// Get возвращает очередной элемент.
// Если элементов нет и block = false -
// возвращает ошибку.
func (q Queue) Get(block bool) (int, error) {
	select {
	case val := <-q:
		return val, nil
	default:
		if block {
			return <-q, nil
		}
		return 0, ErrEmpty
	}
}

// Put помещает элемент в очередь.
// Если очередь заполнения и block = false -
// возвращает ошибку.
func (q Queue) Put(val int, block bool) error {
	select {
	case q <- val:
		return nil
	default:
		if block {
			q <- val
			return nil
		}
		return ErrFull
	}
}

// MakeQueue создает новую очередь
func MakeQueue(n int) Queue {
	return make(chan int, n)
}

// конец решения

func main() {
	q := MakeQueue(2)

	err := q.Put(1, false)
	fmt.Println("put 1:", err)
	fmt.Println("put 1:", err, len(q))

	err = q.Put(2, false)
	fmt.Println("put 2:", err)
	fmt.Println("put 2:", err, len(q))

	err = q.Put(3, false)
	fmt.Println("put 3:", err)
	fmt.Println("put 3:", err, len(q))

	res, err := q.Get(false)
	fmt.Println("get:", res, err)
	fmt.Println("get:", res, err, len(q))

	res, err = q.Get(false)
	fmt.Println("get:", res, err)
	fmt.Println("get:", res, err, len(q))

	res, err = q.Get(false)
	fmt.Println("get:", res, err)
	fmt.Println("get:", res, err, len(q))
}
