package main

import (
	"fmt"
	"runtime"
	"sync"
)

// начало решения

// count отправляет в канал числа от start до бесконечности
func count(cancel chan struct{}, start int) <-chan int {
	out := make(chan int)
	go func() {
		defer func() {
			close(out)
			fmt.Println("count end")
		}()
		for i := start; ; i++ {
			select {
			case out <- i:
			case <-cancel:
				fmt.Println("count cancel")
				return
			}
		}
	}()
	return out
}

// take выбирает первые n чисел из in и отправляет в выходной канал
func take(cancel chan struct{}, in <-chan int, n int) <-chan int {
	out := make(chan int)
	go func() {
		defer func() {
			close(out)
			fmt.Println("take end")
		}()
		for i := 0; i < n; i++ {
			select {
			case out <- <-in:
				fmt.Println("take", i)
			case <-cancel:
				fmt.Println("take cancel")
				return
			}
		}
	}()
	return out
}

// конец решения

func main() {
	defer fmt.Println("defer main: ", runtime.NumGoroutine())
	fmt.Println("main start: ", runtime.NumGoroutine())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		cancel := make(chan struct{})
		defer close(cancel)

		stream := take(cancel, count(cancel, 10), 5)
		fmt.Println("go main: ", runtime.NumGoroutine())
		first := <-stream
		second := <-stream
		third := <-stream

		fmt.Println(first, second, third)
		wg.Done()
	}()
	wg.Wait()
	// time.Sleep(5 * time.Second)
	fmt.Println("after wait: ", runtime.NumGoroutine())
}
