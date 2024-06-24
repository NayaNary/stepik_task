package main

import (
	"fmt"
	"regexp"
	"strings"
)

// начало решения

// slugify возвращает "безопасный" вариант заголовока:
// только латиница, цифры и дефис
func slugify(src string) string {
	src = strings.ToLower(src)
	const validChars = `(\s*[-a-z]+\s*)|(\s*[0-9]+\s*)`
	re := regexp.MustCompile(validChars)
	three := re.FindAllString(src, -1)
	for key, val := range three {
		val = strings.ReplaceAll(val, " ", "")
		three[key] = val
	}
	return strings.Join(three, "-")
}

// конец решения

func main() {
	const phrase = "Go Is Awesome!"
	const want = "go-is-awesome"
	got := slugify(phrase)
	fmt.Println(got, got == want)

}
