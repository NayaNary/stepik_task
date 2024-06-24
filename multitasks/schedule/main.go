package main

import (
	"fmt"
	"time"
)

// начало решения

func schedule(dur time.Duration, fn func()) func() {
	ticker := time.NewTicker(dur)
	canceled := make(chan struct{})

	cancel := func() {
		ticker.Stop()
		select {
		case canceled <- struct{}{}:
		default:
		}

	}
	go func() {
		fmt.Println("schedule start")
		defer func() {
			fmt.Println("schedule end")
			canceled = nil
		}()
		for {
			select {
			case <-ticker.C:
				fn()
			case <-canceled:
				fmt.Println("schedule canceled")
				return
			}
		}
	}()
	fmt.Println("schedule", canceled)
	return cancel
}

// конец решения

func main() {
	work := func() {
		at := time.Now()
		fmt.Printf("%s: work done\n", at.Format("15:04:05.000"))
	}

	cancel := schedule(50*time.Millisecond, work)
	cancel()
	cancel()
	defer cancel()

	// хватит на 5 тиков
	time.Sleep(260 * time.Millisecond)
	cancel()
	time.Sleep(260 * time.Millisecond)
	cancel()

}
