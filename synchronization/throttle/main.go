// Ограничитель вызовов
package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrBusy = errors.New("busy")
var ErrCanceled = errors.New("canceled")

// начало решения

// throttle следит, чтобы функция fn выполнялась не более limit раз в секунду.
// Возвращает функции handle (выполняет fn с учетом лимита) и cancel (останавливает ограничитель).
func throttle(limit int, fn func()) (handle func() error, cancel func()) {
	canceled := make(chan struct{}, limit)
	mx := sync.Mutex{}
	counterLimit := 0
	// isCanceled := false
	tiker := time.NewTicker(time.Second)
	handle = func() error {
		select {
		case <-tiker.C:
			mx.Lock()
			counterLimit = 0
			fmt.Println("<-tiker.C", counterLimit)
			mx.Unlock()
		case _, ok := <-canceled:
			if !ok {
				return ErrCanceled
			}
			tiker.Stop()
			fmt.Println("ErrCanceled", ErrCanceled)
			return ErrCanceled
		default:
		}
		if counterLimit == limit {
			fmt.Println("ErrBusy", ErrBusy)
			return ErrBusy
		}
		// fmt.Println("isCanceled", isCanceled)
		// if isCanceled {
		// 	return ErrCanceled
		// }
		mx.Lock()
		counterLimit += 1
		mx.Unlock()
		fn()
		return nil
	}

	cancel = func() {
		// if isCanceled {
		// 	return
		// }
		select {
		case <-canceled:
			fmt.Println("<-canceled")
		default:
			mx.Lock()
			close(canceled)
			// isCanceled = true
			mx.Unlock()
			fmt.Println("close(canceled)")
		}
	}

	return
}

// конец решения

func main() {
	work := func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}

	handle, cancel := throttle(10, work)
	defer cancel()

	const n = 10
	var nOK, nErr int
	for i := 0; i < n; i++ {
		err := handle()
		if err == nil {
			nOK += 1
		} else {
			nErr += 1
		}
	}
	fmt.Println()
	fmt.Printf("%d calls: %d OK, %d busy\n", n, nOK, nErr)

	// work := func() {
	// 	time.Sleep(200 * time.Millisecond)
	// 	fmt.Print(".")
	// }

	// handle, cancel := throttle(10, work)
	// defer cancel()

	// const n = 10
	// var nOK, nErr int
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// t := time.Now()
	// go func() {
	// 	defer wg.Done()
	// 	time.Sleep(500 * time.Millisecond)
	// 	cancel()
	// 	cancel()
	// }()
	// for i := 0; i < n; i++ {
	// 	err := handle()
	// 	if err == nil {
	// 		nOK += 1
	// 	} else {
	// 		nErr += 1
	// 	}
	// }

	// wg.Wait()
	// fmt.Println("time: ", time.Since(t))
	// fmt.Println()
	// fmt.Printf("%d calls: %d OK, %d busy\n", n, nOK, nErr)

	// work := func() {
	// 	time.Sleep(2000 * time.Millisecond)
	// 	fmt.Print(".")
	// }

	// handle, cancel := throttle(10, work)
	// defer cancel()

	// const n = 10
	// var nOK, nErr int
	// mx := sync.Mutex{}
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	time.Sleep(4000 * time.Millisecond)
	// 	cancel()
	// }()
	// for i := 0; i < n; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		defer wg.Done()
	// 		err := handle()
	// 		if err == nil {
	// 			mx.Lock()
	// 			nOK += 1
	// 			mx.Unlock()
	// 		} else {
	// 			mx.Lock()
	// 			nErr += 1
	// 			mx.Unlock()
	// 		}
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println()
	// fmt.Printf("%d calls: %d OK, %d busy\n", n, nOK, nErr)
}
