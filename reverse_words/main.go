package main

import (
	"fmt"
	"math/rand"
	_ "net/http/pprof"
	"runtime"
	"strings"
	"sync"
)

// начало решения

// генерит случайные слова из 5 букв
// с помощью randomWord(5)
func generate(cancel <-chan struct{}) <-chan string {
	out := make(chan string)
	go func() {
		defer func() {
			close(out)
			fmt.Println("generate end")
		}()
		for {
			select {
			case out <- randomWord(5):
			case <-cancel:
				fmt.Println("generate cancel")
				return
			}
		}
	}()
	return out
}

// выбирает слова, в которых не повторяются буквы,
// abcde - подходит
// abcda - не подходит
func takeUnique(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer func() {
			close(out)
			fmt.Println("takeUnique end")
		}()
		for {
			select {
			case word, ok := <-in:
				if !ok {
					return
				}
				if ok := uniqueChars(word); ok {
					select {
					case <-cancel:
						fmt.Println("takeUnique cancel")
						return
					case out <- word:
					}
				}
			case <-cancel:
				fmt.Println("takeUnique cancel")
				return
			}
		}
	}()
	return out
}
func uniqueChars(world string) bool {
	for _, char := range world {
		if res := strings.Count(world, string(char)); res > 1 {
			return false
		}
	}
	return true
}

// переворачивает слова
// abcde -> edcba
func reverse(cancel <-chan struct{}, in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer func ()  {
			close(out)
			fmt.Println("reverse end")
		}()
		for {
			select {
			case out <- reverseChars(<-in):
			case <-cancel:
				fmt.Println("reverse cancel")
				return
			}
		}
	}()
	return out
}

func reverseChars(word string) (reverseword string) {
	reverseword = word + " -> "
	for i := len(word) - 1; i >= 0; i-- {
		reverseword += string(word[i])
	}
	return
}

// объединяет c1 и c2 в общий канал
func merge(cancel <-chan struct{}, c1, c2 <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer func(){
			close(out)
			fmt.Println("merge end")
		}() 
		for c1 != nil || c2 != nil {
			select {
			case val1, ok := <-c1:
				if ok {
					select {
					case <-cancel:
						return
					case out <- val1:
					}
				} else {
					c1 = nil
				}

			case val2, ok := <-c2:
				if ok {
					select {
					case <-cancel:
						return
					case out <- val2:
					}
				} else {
					c2 = nil
				}
			case <-cancel:
				fmt.Println("merge cancel")
				return
			}
		}
	}()
	return out
}

// печатает первые n результатов
func print(cancel <-chan struct{}, in <-chan string, n int) {
	for i := 0; i < n; i++ {
		select {
		case word := <-in:
			fmt.Println(word)
		case <-cancel:
			return
		}
	}
}

// конец решения

// генерит случайное слово из n букв
func randomWord(n int) string {
	const letters = "aeiourtnsl"
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = letters[rand.Intn(len(letters))]
	}
	return string(chars)
}

func main() {
	defer fmt.Println("defer main: ", runtime.NumGoroutine())
	wg := sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("main start: ", runtime.NumGoroutine())
	go func() {
		cancel := make(chan struct{})
		defer close(cancel)

		c1 := generate(cancel)
		c2 := takeUnique(cancel, c1)
		c3_1 := reverse(cancel, c2)
		c3_2 := reverse(cancel, c2)
		c4 := merge(cancel, c3_1, c3_2)
		print(cancel, c4, 10)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("main end: ", runtime.NumGoroutine())

}
