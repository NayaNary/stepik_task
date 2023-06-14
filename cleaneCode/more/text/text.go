package text

import (
	"more/numbers"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// ArePermutation checks if two strings are permutations of each other.
func ArePermutation(first, second string) bool {
	if len(first) != len(second) {
		return false
	}
	runes1st := []rune(first)
	runes2nd := []rune(second)
	sort.Slice(runes1st, func(i, j int) bool { return runes1st[i] < runes1st[j] })
	sort.Slice(runes2nd, func(i, j int) bool { return runes2nd[i] < runes2nd[j] })
	return reflect.DeepEqual(runes1st, runes2nd)
}

// AsRunes returns the number as a slice of it's digits runes.
func AsRunes(n int) []rune {
	return []rune(strconv.Itoa(n))
}

// AsDigitString returns the number as a string composed of it's digits,
// separated by dashes. E.g. 42513 → 4-2-5-1-3
func AsDigitString(n int) string {
	digits := numbers.AsDigits(n)
	parts := make([]string, len(digits))
	for idx, d := range digits {
		parts[idx] = strconv.Itoa(d)
	}
	return strings.Join(parts, "-")
}
