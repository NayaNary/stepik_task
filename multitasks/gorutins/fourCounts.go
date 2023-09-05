package gorutins

import "fmt"

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWordsFourCount(next nextFunc) counter {
	pending := make(chan string)
	go submitWords(next, pending)

	done := make(chan struct{})
	counted := make(chan pair)

	// начало решения

	// запустите четыре горутины countWords()
	// вместо одной
	for i := 1; i <= 4; i++ {
		go countWordsFourCount(i, done, pending, counted)
	}

	// используйте канал завершения, чтобы дождаться
	// окончания обработки и закрыть канал counted
	go func() {
		for i := 0; i < 4; i++ { // (4)
			<-done

		}
		close(counted)
	}()
	// конец решения

	return fillStats(counted)
}

// countWords считает цифры в словах
func countWordsFourCount(id int, done chan<- struct{}, in <-chan string, out chan<- pair) {
	fmt.Println("нопер горутины: ", id)
	for word := range in {
		count := countDigits(word)
		out <- pair{
			word:  word,
			count: count,
		}
	}
	done <- struct{}{}
}

func CountDigitsInWordsFourCount() {
	phrase := "1 22 333 4444 55555 666666 7777777 88888888"
	next := wordGenerator(phrase)
	stats := countDigitsInWordsFourCount(next)
	printStats(stats)
}
