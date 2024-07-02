package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"time"
)

// начало решения

// ErrFailed и ErrManual - причины остановки цикла.
var ErrFailed = errors.New("failed")
var ErrManual = errors.New("manual")

// Worker выполняет заданную функцию в цикле, пока не будет остановлен.
// Гарантируется, что Worker используется только в одной горутине.
type Worker struct {
	fn          func() error
	ctx         context.Context
	cancel      context.CancelCauseFunc
	isStart     bool
	isStop      bool
}

// NewWorker создает новый экземпляр Worker с заданной функцией.
// Но пока не запускает цикл с функцией.
func NewWorker(fn func() error) *Worker {
	ctx, cancel := context.WithCancelCause(context.Background())
	return &Worker{
		fn:     fn,
		ctx:    ctx,
		cancel: cancel,
	}
}

// Start запускает отдельную горутину, в которой циклически
// выполняет заданную функцию, пока не будет вызван метод Stop,
// либо пока функция не вернет ошибку.
// Повторные вызовы Start игнорируются.
func (w *Worker) Start() {
	if w.isStart {
		return
	}
	w.isStart = true
	// count := 0
	go func() {
		defer func() {
			w.isStart = false
		}()
		for {
			select {
			case <-w.ctx.Done():
				return
			default:
				err := w.fn()
				if err != nil {
					w.cancel(ErrFailed)
					return
				}
			}
		}
	}()
}

// Stop останавливает выполнение цикла.
// Вызов Stop до Start игнорируется.
// Повторные вызовы Stop игнорируются.
func (w *Worker) Stop() {
	// TODO: реализовать требования
	if !w.isStart || w.isStop {
		return
	}

	w.isStop = true
	w.cancel(ErrManual)
}

// AfterStop регистрирует функцию, которая
// будет вызвана после остановки цикла.
// Можно зарегистрировать несколько функций.
// Вызовы AfterStop после Start игнорируются.
func (w *Worker) AfterStop(fn func()) {
	// TODO: реализовать требования
	if w.isStart {
		return
	}
	context.AfterFunc(w.ctx, fn)
}

// Err возвращает причину остановки цикла:
// - ErrManual - вручную через метод Stop;
// - ErrFailed - из-за ошибки, которую вернула функция.
func (w *Worker) Err() error {
	return context.Cause(w.ctx)
}

// конец решения

func main() {
	{
		// Start-Stop
		count := 10
		fn := func() error {
			fmt.Print(count, " ")
			count--
			time.Sleep(10 * time.Millisecond)
			return nil
		}

		worker := NewWorker(fn)
		fmt.Println(runtime.NumGoroutine())
		worker.Start()
		time.Sleep(20 * time.Millisecond)
		worker.Stop()
		// fmt.Println(runtime.NumGoroutine())
		// worker.Stop()
		// fmt.Println(runtime.NumGoroutine())
		// worker.Stop()
		// fmt.Println(runtime.NumGoroutine())

		fmt.Println("Err:", worker.Err())
		fmt.Println(runtime.NumGoroutine())
		fmt.Println()
		// 10 9 8 7 6 5 4 3 2 1
	}
	// {
	// 	// Start-Stop
	// 	count := 10
	// 	fn := func() error {
	// 		fmt.Print(count, " ")
	// 		count--
	// 		time.Sleep(10 * time.Millisecond)
	// 		return nil
	// 	}

	// 	worker := NewWorker(fn)
	// 	worker.Start()
	// 	time.Sleep(105 * time.Millisecond)
	// 	worker.Stop()

	// 	fmt.Println("Err:", worker.Err())

	// 	fmt.Println()
	// 	// 10 9 8 7 6 5 4 3 2 1
	// }
	// {
	// 	// ErrFailed
	// 	count := 3
	// 	fn := func() error {
	// 		fmt.Print(count, " ")
	// 		count--
	// 		if count == 0 {
	// 			return errors.New("count is zero")
	// 		}
	// 		time.Sleep(10 * time.Millisecond)
	// 		return nil
	// 	}

	// 	worker := NewWorker(fn)
	// 	worker.Start()
	// 	time.Sleep(35 * time.Millisecond)
	// 	worker.Stop()

	// 	fmt.Println(worker.Err())
	// 	// 3 2 1 failed
	// }
	// {
	// 	// AfterStop
	// 	fn := func() error { return nil }

	// 	worker := NewWorker(fn)
	// 	worker.AfterStop(func() {
	// 		fmt.Println("called after stop")
	// 	})

	// 	worker.Start()
	// 	worker.Stop()

	// 	time.Sleep(10 * time.Millisecond)
	// 	// called after stop
	// }
}
