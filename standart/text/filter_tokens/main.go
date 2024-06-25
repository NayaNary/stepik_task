package main

import (
	"fmt"
	"io"
	"strings"
)

// TokenReader начитывает токены из источника
type TokenReader interface {
	// ReadToken считывает очередной токен
	// Если токенов больше нет, возвращает ошибку io.EOF
	ReadToken() (string, error)
}

// TokenWriter записывает токены в приемник
type TokenWriter interface {
	// WriteToken записывает очередной токен
	WriteToken(s string) error
	Words() []string
}

// начало решения
type tokenReader struct {
	bf []string
}

// ReadToken считывает очередной токен
// Если токенов больше нет, возвращает ошибку io.EOF
func (w *tokenReader) ReadToken() (string, error) {
	if len(w.bf) == 0 {
		return "", io.EOF
	}
	token := w.bf[0]
	w.bf = w.bf[1:]
	return token, nil
}

func NewWordReader(data string) TokenReader {
	return &tokenReader{bf: strings.Split(data, " ")}
}

type tokenWriter struct {
	bf []string
}

// WriteToken записывает очередной токен
func (w *tokenWriter) WriteToken(s string) error {
	w.bf = append(w.bf, s)
	return nil
}

func (w *tokenWriter) Words() []string {
	return w.bf
}

func NewWordWriter() TokenWriter {
	return &tokenWriter{bf: make([]string, 0)}
}

// FilterTokens читает все токены из src и записывает в dst тех,
// кто проходит проверку predicate
func FilterTokens(dst TokenWriter, src TokenReader, predicate func(s string) bool) (int, error) {
	amountToken := 0
	for {
		token, err := src.ReadToken()
		if err == io.EOF {
			return amountToken, nil
		}
		if err != nil {
			return amountToken, err
		}
		if predicate(token) {
			if err := dst.WriteToken(token); err != nil {
				return amountToken, err
			}
			amountToken++
		}
	}
}

// конец решения

func main() {
	// Для проверки придется создать конкретные типы,
	// которые реализуют интерфейсы TokenReader и TokenWriter.

	// Ниже для примера используются NewWordReader и NewWordWriter,
	// но вы можете сделать любые на свое усмотрение.

	r := NewWordReader("go is awesome")
	w := NewWordWriter()
	predicate := func(s string) bool {
		return s != "is"
	}
	n, err := FilterTokens(w, r, predicate)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d tokens: %v\n", n, w.Words())
	// 2 tokens: [go awesome]
}
