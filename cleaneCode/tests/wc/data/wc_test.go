package data

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func BenchmarkRegexp(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		rand.Seed(0)
		phrase := randomPhrase(length)
		name := fmt.Sprintf("Regexp-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				WordCountRegexp(phrase)
			}
		})
	}
}

func BenchmarkFields(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		rand.Seed(0)
		phrase := randomPhrase(length)
		name := fmt.Sprintf("Fields-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				WordCountFields(phrase)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		rand.Seed(0)
		phrase := randomPhrase(length)
		name := fmt.Sprintf("Split-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				WordCountSplit(phrase)
			}
		})
	}
}

func BenchmarkAllocate(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		rand.Seed(0)
		phrase := randomPhrase(length)
		name := fmt.Sprintf("Allocate-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				WordCountAllocate(phrase)
			}
		})
	}
}


func BenchmarkLowerPhrase(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		rand.Seed(0)
		phrase := randomPhrase(length)
		name := fmt.Sprintf("LowerPhrase-%d", length)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				WordCountLowerPhrase(phrase)
			}
		})
	}
}


func BenchmarkIndex(b *testing.B) {
	for _, length := range []int{10, 100, 1000, 10000} {
		rand.Seed(0)
		phrase := randomPhrase(length)
		name := fmt.Sprintf("Index-%d", length)
		srt := MakeWords(phrase)
		words := strings.Split(phrase, " ")
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				srt.Index(words[0])
			}
		})
	}
}

// randomPhrase returns a phrase of n random words
func randomPhrase(n int) string {
	set := ""
	for i := 0; i < n; i++ {
		n := rand.Intn(100000)
		set += fmt.Sprintf(" %d", n)
	}
	return set
}
