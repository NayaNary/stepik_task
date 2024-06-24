package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

// начало решения

func delay(dur time.Duration, fn func()) func() {
	timer := time.NewTimer(dur)
	cancelRes := make(chan struct{})
	cancel := func() {
		if cancelRes != nil {
			cancelRes <- struct{}{}
		}
	}
	go func() {
		defer func() {
			cancelRes = nil
		}()
		select {
		case <-timer.C:
			fn()
			return
		case <-cancelRes:
			return
		}
	}()
	return cancel
}

// конец решения

func main() {
	defer fmt.Println("defer main: ", runtime.NumGoroutine())
	fmt.Println("main start: ", runtime.NumGoroutine())
	rand.Seed(time.Now().Unix())

	work := func() {
		fmt.Println("work done")
	}

	cancel := delay(100*time.Millisecond, work)

	cancel()

	// cancel()

	time.Sleep(10 * time.Millisecond)
	if rand.Float32() < 0.5 {
		cancel()
		fmt.Println("delayed function canceled")
	}
	time.Sleep(100 * time.Millisecond)

	cancel()
	fmt.Println("after wait: ", runtime.NumGoroutine())
}
