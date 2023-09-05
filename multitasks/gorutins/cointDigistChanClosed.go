package gorutins

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWordsChanClose(next nextFunc) counter {
	pending := make(chan string)
	go submitWordsChanClose(next, pending)

	counted := make(chan pair)
	go countWordsChanClose(pending, counted)

	return fillStatsChanClose(counted)
}

// начало решения

// submitWords отправляет слова на подсчет
func submitWordsChanClose(next nextFunc, out chan string) {
	for {
		word := next()
		out <- word
		if word == "" {
			break
		}
	}
	close(out)
}

// countWords считает цифры в словах
func countWordsChanClose(in chan string, out chan pair) {
	for word := range in {
		count := countDigits(word)
		out <- pair{
			word:  word,
			count: count,
		}
	}
	close(out)
}

// fillStats готовит итоговую статистику
func fillStatsChanClose(in chan pair) counter {
	stats := counter{}
	for info := range in {
		if info.word == "" {
			continue
		}
		stats[info.word] = info.count
	}
	return stats
}

// конец решения

func CountDigitsInWordsChanClose() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWordsChanClose(next)
	printStats(stats)
}
