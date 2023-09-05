package data

import (
	"strings"
)

// Words represents words in a string.
type Words struct {
	mapWords map[string]int
}

// MakeWords creates words from a string.
func MakeWords(s string) Words {
	words := strings.Fields(s)
	size := len(words) / 2
	if size > 10000 {
		size = 10000
	}
	mapWords := make(map[string]int, size)
	for index, word := range words {
		if _, ok := mapWords[word]; !ok {
			mapWords[word] = index
		}
	}
	return Words{mapWords: mapWords}
}

// Index returns the index of the first instance of word in words,
// or -1 if word is not present in words.
func (w Words) Index(word string) int {
	if index, ok := w.mapWords[word]; ok {
		return index
	}
	return -1
}
