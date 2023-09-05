package gorutins

// countDigitsInWords считает количество цифр в словах,
// выбирая очередные слова с помощью next()
func countDigitsInWordsTwoGo(next nextFunc) counter {
	pending := make(chan string)
	counted := make(chan pair)
	stats := counter{}

	// начало решения

	// отправляет слова на подсчет
	go func() {
		// Пройдите по словам и отправьте их
		// в канал pending
		for {
			word := next()
			pending <- word
			if word == "" {
				break
			}
		}
	}()

	// считает цифры в словах
	go func() {
		// Считайте слова из канала pending,
		// посчитайте количество цифр в каждом,
		// и запишите его в канал counted
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

func CountDigitsInWordsTwoGo() {
	phrase := "0ne 1wo thr33 4068"
	next := wordGenerator(phrase)
	stats := countDigitsInWordsTwoGo(next)
	printStats(stats)
}
