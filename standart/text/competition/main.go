package main

import (
	"fmt"
	"regexp"
	"strings"
)

var wordRE = regexp.MustCompile(`[a-z0-9\-]+`)

// начало решения

func slugifyRegExp(src string) string {
	words := wordRE.FindAllString(strings.ToLower(src), -1)
	return strings.Join(words, "-")
}

func slugify(src string) string {
	res := strings.ToLower(src)
	res = strings.Map(purifyChar, res)
	words := strings.Fields(res)
	return strings.Join(words, "-")
}
func purifyChar(r rune) rune {
	const validChars string = "abcdefghijklmnopqrstuvwxyz01234567890- "
	if strings.IndexRune(validChars, r) == -1 {
		return ' '
	}
	return r
}

func slugifyMy(src string) string {
	const validChars = "qQwWeErRtTyYuUiIoOpPAaSsDdFfGghHjJkKlLzZxXcCvVbBnNmM0147852369-"
	for _, char := range src {
		if !strings.Contains(validChars, string(char)) {
			src = strings.ReplaceAll(src, string(char), " ")
		}
	}
	// перевод в нижний регистр
	src = strings.ToLower(src)
	// разделить на слова
	masFields := strings.Fields(src)
	// объединить в стоку, через -
	return strings.Join(masFields, "-")
}

func slugifyMyNew(src string) string {
	res := []byte(src)
	for i := 0; i < len(res); i++ {
		if 'A' <= res[i] && res[i] <= 'Z' {
			res[i] = res[i] + 32
		}
	}
	for i := 0; i < len(res); i++ {
		if !isGood(res[i]) {
			j := i
			for ; j < len(res); j++ {
				if isGood(res[j]) {
					if i != 0 {
						res[i] = '-'
						i++
					}
					res = append(res[:i], res[j:]...)
					del := j-i
					i = j
					if del == 0 {					
						i--
					} else {
						i = i - (del+1)
					}
					break
				}
			}
			if j == len(res) {
				res = res[:i]
			}
		}
	}

	return string(res)
}

func isGood(char byte) bool {
	return ('a' <= char && char <= 'z') || ('0' <= char && char <= '9') || char == '-'
}

func test(src string) string {
	srsBytes := []byte(src)
	for idx, char := range srsBytes {
		if int(char) >= 65 && int(char) <= 90 {
			srsBytes[idx] = uint8(int(char) + 32)
		}
		if !((int(char) >= 30 && int(char) <= 57) && (int(char) >= 97 && int(char) <= 122)) {
			srsBytes[idx] = uint8(45)
		}
	}
	return string(srsBytes)
}

// конец решения

func main() {
	const phrase = "A 100x Investment (2019)"
	const want = "a-100x-investment-2019"
	// const phrase = "Hello, 中国!"
	// const want = "hello"
	// const phrase = "Go Talks: Cuddle: an App Engine Demo"
	// const want = "go-talks-cuddle-an-app-engine-demo"
	// const phrase = "Go at Google I/O"
	// const want = "go-at-google-i-o"
	// const phrase = "!Attention, attention!"
	// const want = "attention-attention"
	// const phrase = "Understanding the x64 code models (2012)"
	// const want = "understanding-the-x64-code-models-2012"
	// const phrase = "Go - Is - Awesome"
	// const want = "go---is---awesome"
	// const phrase = "Кто я!!"
	// const want = ""
	// const phrase = "Debugging Go code (a status report)"
	// const want = "debugging-go-code-a-status-report"
	// const phrase = "We haven’t killed 90% of all plankton"
	// const want = "we-haven-t-killed-90-of-all-plankton"

	slug := slugifyMyNew(phrase)
	fmt.Println(phrase)
	fmt.Println(slug)
	fmt.Println(slug == want)

}
