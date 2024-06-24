package numbers

import (
    "more/text"
)

// IsEven checks if the number is even.
func IsEven(n int) bool {
	return n%2 == 0
}

// AsDigits returns a slice of digits that make up the number.
func AsDigits(n int) []int {
	runes := text.AsRunes(n)
	count := len(runes)
	zero := int('0')
	digits := make([]int, count)
	for idx, char := range runes {
		digits[idx] = int(char) - zero
	}
	return digits
}
