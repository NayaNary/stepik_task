package main

import (
	"errors"
	"fmt"
	"time"
)

var ErrCanceled error = errors.New("canceled")

// начало решения
func withRateLimit(limit int, fn func()) (handle func() error, cancel func()) {
	canceled := make(chan struct{})
	duration := time.Duration(int64(time.Second) / int64(limit))
	ticker := time.NewTicker(duration)
	handle = func() error {
		select {
		case <-ticker.C:
			go func() {
				fn()
			}()
		case <-canceled:
			return ErrCanceled
		}
		return nil
	}

	cancel = func() {
		select {
		case <-canceled:
			return
		default:
			ticker.Stop()
			close(canceled)
		}
	}

	return
}

// конец решения

func main() {
	count := 1
	work := func() {
		fmt.Println(count, ".")
		count++
		time.Sleep(time.Millisecond * 500)
	}

	handle, cancel := withRateLimit(10, work)
	defer cancel()
	start := time.Now()
	const n = 10
	for i := 0; i < n; i++ {
		if err := handle(); err != nil {
			fmt.Println(err)
		}

	}
	fmt.Println()
	fmt.Printf("%d queries took %v\n", n, time.Since(start))
	// tikerStart := time.NewTicker(time.Duration(int(time.Second) / 5))
	// tikerEnd := time.NewTicker(time.Duration(5 * time.Second))
	// start := time.Now()

	// handle := func(i int) error {
	// 	select {
	// 	case <-tikerStart.C:
	// 		fmt.Println("i", i)
	// 	case val := <-tikerEnd.C:
	// 		fmt.Println("i", i, "time:", val)
	// 		return nil
	// 	}
	// 	return nil
	// }
	// for i := 0; i < 10; i++ {
	// 	handle(i)
	// }
	// fmt.Println()
	// fmt.Printf("%d queries took %v\n", 10, time.Since(start))
}
