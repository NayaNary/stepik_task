package main

import (
	"fmt"
	"strings"
)

// начало решения

// slugify возвращает "безопасный" вариант заголовока:
// только латиница, цифры и дефис
func slugify(src string) string {
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

// конец решения

func main() {
	const phrase = "Why Generics?"
	const want = "why-generics"
	// const phrase = "JSON-RPC: a tale of interfaces"
	// const want = "json-rpc-a-tale-of-interfaces"
	// const phrase = "Go at Google I/O"
	// const want = "go-at-google-i-o"
	// const p = "Hello, 中国!"
	// const pw ="hello"
	got := slugify(phrase)
	fmt.Println(got, got == want)
	// fmt.Println(got, got==pw)
	// if got != want {
	// 	fmt.Printf("%s: got %#v, want %#v", phrase, got, want)
	// }
}
