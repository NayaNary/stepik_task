package gorutins

import (
	"strings"
)

// nextFunc возвращает следующее слово из генератора
type nextFunc func() string

// pair хранит слово и количество цифр в нем
type pair struct {
	word  string
	count int
}

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWordsGen(next nextFunc) counter {
	counted := make(chan pair)
	stats := counter{}

	// начало решения

	go func() {
		for {
			word := next()

			// посчитайте количество цифр в каждом,
			count := countDigits(word)
			wordInfo := pair{
				word:  word,
				count: count,
			}
			counted <- wordInfo
			if word == "" {
				break
			}
		}

		// и запишите его в канал counted
	}()

	// Считайте значения из канала counted
	// и заполните stats.
	for {
		info := <-counted
		if info.word == "" {
			break
		}
		stats[info.word] = info.count
	}

	// В результате stats должна содержать слова
	// и количество цифр в каждом.

	// конец решения

	return stats
}

// wordGenerator возвращает генератор, который выдает слова из фразы
func wordGenerator(phrase string) nextFunc {
	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}
}

func CountDigitsInWordsGen() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWordsGen(next)
	printStats(stats)
}
