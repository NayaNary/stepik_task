package gorutins

import (
	"strings"
)

// countDigitsInWords считает количество цифр в словах фразы
func countDigitsInWordsChan(phrase string) counter {
	words := strings.Fields(phrase)
	counted := make(chan int)
	stats := counter{}

	// начало решения

	go func() {
		for _, word := range words {
			counted <- countDigits(word)
		}
	}()

	for _, word := range words {
		stats[word] = <-counted
	}

	return stats
}

func CountDigitsInWordsChan() {
	phrase := "0ne 1wo thr33 4068"
	stats := countDigitsInWords(phrase)
	printStats(stats)
}
