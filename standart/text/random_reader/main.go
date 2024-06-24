package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
)

type randomReader struct{}

func (r *randomReader) Read(p []byte) (n int, err error) {
    return rand.Read(p)
}

func RandomReaderEx(max int) io.Reader {
    rd := &randomReader{}
    return io.LimitReader(rd, int64(max))
}

// начало решения

// RandomReader создает читателя, который возвращает случайные байты,
// но не более max штук
func RandomReader(max int) io.Reader {
	token := make([]byte, max)
	rand.Read(token)
	reader := bytes.NewReader(token)
	return reader
}

// конец решения

func main() {
	rand.Seed(0)

	rnd := RandomReader(5)
	rd := bufio.NewReader(rnd)
	for {
		b, err := rd.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d ", b)
	}
	fmt.Println()
	// 1 148 253 194 250
}
