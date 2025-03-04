// Пишем рандеву
package main

import (
	"fmt"
	"sync"
	"time"
)

// начало решения

// Rendezvous представляет рандеву двух горутин.
type Rendezvous struct {
	goChanOne chan struct{}
	goChanTwo chan struct{}
	mx        sync.Mutex
	countCall uint
}

// NewRendezvous создает новое рандеву.
func NewRendezvous() *Rendezvous {
	return &Rendezvous{
		goChanOne: make(chan struct{}),
		goChanTwo: make(chan struct{}),
		mx:        sync.Mutex{},
	}
}

// Ready фиксирует, что вызывающая горутина прибыла к точке сбора.
// Блокирует вызывающую горутину, пока не прибудет вторая.
// Когда обе горутины прибудут, Ready их разблокирует.
func (r *Rendezvous) Ready() {
	wg := sync.WaitGroup{}

	if r.countCall == 0 {
		r.mx.Lock()
		r.countCall++
		r.mx.Unlock()
		wg.Add(1)
		go func() {
			defer wg.Done()
			close(r.goChanOne)
			<-r.goChanTwo
		}()
	}

	if r.countCall == 1 {
		r.mx.Lock()
		r.countCall++
		r.mx.Unlock()
		wg.Add(1)
		go func() {
			defer wg.Done()
			close(r.goChanTwo)
			<-r.goChanOne
		}()
	}

	wg.Wait()
}

// конец решения

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	rend := NewRendezvous()

	go func() {
		fmt.Println("1: started")
		time.Sleep(10 * time.Millisecond)
		fmt.Println("1: reached the meeting point")

		rend.Ready()

		fmt.Println("1: going further")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("1: done")
		wg.Done()
	}()

	time.Sleep(20 * time.Millisecond)

	go func() {
		fmt.Println("2: started")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("2: reached the meeting point")

		rend.Ready()

		fmt.Println("2: going further")
		time.Sleep(10 * time.Millisecond)
		fmt.Println("2: done")
		wg.Done()
	}()

	wg.Wait()

	/*
		1: started
		1: reached the meeting point
		2: started
		2: reached the meeting point
		2: going further
		1: going further
		2: done
		1: done
	*/
}
