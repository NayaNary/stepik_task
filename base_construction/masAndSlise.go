package baseconstruction

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func CutString() {
	var text, res string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...

	if len(text) > width {
		res = text[:width]
		res += "..."
	} else {
		res = text
	}

	fmt.Println(res)

}

// printCounter печатает карту в формате
// key1:val1 key2:val2 ... keyN:valN
func printCounter(counter map[int]int) {
	digits := make([]int, 0)
	for d := range counter {
		digits = append(digits, d)
	}
	sort.Ints(digits)
	for idx, digit := range digits {
		fmt.Printf("%d:%d", digit, counter[digit])
		if idx < len(digits)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}

func CountNums() {
	var number int
	fmt.Scanf("%d", &number)

	// Посчитайте, сколько раз каждая цифра встречается
	// в числе `number`. Запишите результат в карту `counter`,
	// где ключом является цифра числа, а значением -
	// сколько раз она встречается
	counter := make(map[int]int)
	// ...
	for {
		num := number % 10
		number = number / 10
		count, ok := counter[num]
		if ok {
			counter[num] = count + 1
		} else {
			counter[num] = 1
		}

		if number == 0 {
			break
		}
	}
	printCounter(counter)
}

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}

func Abbr() {
	phrase := readString()

	// 1. Разбейте фразу на слова, используя `strings.Fields()`
	// 2. Возьмите первую букву каждого слова и приведите
	//    ее к верхнему регистру через `unicode.ToUpper()`
	// 3. Если слово начинается не с буквы, игнорируйте его
	//    проверяйте через `unicode.IsLetter()`
	// 4. Составьте слово из получившихся букв и запишите его
	//    в переменную `abbr`
	// ...

	fields := strings.Fields(phrase)
	abbr := []rune{}
	for _, field := range fields {
		if len(field) > 0 {
			for _, char := range field {
				if unicode.IsLetter(rune(char)) {
					abbr = append(abbr, unicode.ToUpper(char))
				}
				break
			}
		}
	}
	fmt.Println(string(abbr))
}