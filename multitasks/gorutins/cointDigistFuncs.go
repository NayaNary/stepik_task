package gorutins

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWordsFuncs(next nextFunc) counter {
	pending := make(chan string)
	go submitWords(next, pending)

	counted := make(chan pair)
	go countWords(pending, counted)

	return fillStats(counted)
}

// начало решения

// submitWords отправляет слова на подсчет
func submitWords(next nextFunc, pending chan string) {
	for {
		word := next()
		pending <- word
		if word == "" {
			break
		}
	}
}

// countWords считает цифры в словах
func countWords(pending chan string, counted chan pair) {
	for {
		word := <-pending
		count := countDigits(word)
		counted <- pair{
			word:  word,
			count: count,
		}
		if word == "" {
			break
		}
	}
}

// fillStats готовит итоговую статистику
func fillStats(counted chan pair) counter {
	stats := counter{}
	for {
		info := <-counted
		if info.word == "" {
			break
		}
		stats[info.word] = info.count

	}
	return stats
}

// конец решения

func CountDigitsInWordsFuncs() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWordsFuncs(next)
	printStats(stats)
}
