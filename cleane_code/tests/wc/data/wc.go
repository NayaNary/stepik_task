package data

import (
	"regexp"
	"strings"
)

// Counter maps words to their counts
type Counter map[string]int

var splitter *regexp.Regexp = regexp.MustCompile(" ")

// WordCountRegexp counts absolute frequencies of words in a string.
// Uses Regexp.Split() to split the string into words.
func WordCountRegexp(s string) Counter {
	counter := make(Counter)
	for _, word := range splitter.Split(s, -1) {
		word = strings.ToLower(word)
		counter[word]++
	}
	return counter
}

// WordCountFields counts absolute frequencies of words in a string.
// Uses strings.Fields() to split the string into words.
func WordCountFields(s string) Counter {
    counter := make(Counter)
    for _, word := range strings.Fields(s) {
        word = strings.ToLower(word)
        counter[word]++
    }
    return counter
}

// WordCountSplit counts absolute frequencies of words in a string.
// Uses strings.Split() to split the string into words.
func WordCountSplit(s string) Counter {
    counter := make(Counter)
    for _, word := range strings.Split(s, " ") {
        word = strings.ToLower(word)
        counter[word]++
    }
    return counter
}

// WordCountLowerPhrase counts absolute frequencies of words in a string.
// Converts the whole string to lower case before splitting.
func WordCountLowerPhrase(s string) Counter {
    counter := make(Counter)
    phrase := strings.ToLower(s)
    for _, word := range strings.Split(phrase, " ") {
        counter[word]++
    }
    return counter
}

// WordCountAllocate counts absolute frequencies of words in a string.
// Pre-allocates memory for the counter.
func WordCountAllocate(s string) Counter {
    words := strings.Split(s, " ")
    size := len(words) / 2
    if size > 10000 {
        size = 10000
    }
    counter := make(Counter, size)
    for _, word := range words {
        word = strings.ToLower(word)
        counter[word]++
    }
    return counter
}