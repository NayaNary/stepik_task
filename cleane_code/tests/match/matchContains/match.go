package matchcontains

import (
	"regexp"
	"strings"
)

// MatchContains returns true if the string
// contains the pattern, false otherwise.
func MatchContains(pattern string, src string) bool {
	return strings.Contains(src, pattern)
}

// Самописная
func MatchContainsCustom(pattern string, src string) bool {
    if pattern == "" {
        return true
    }
    if len(pattern) > len(src) {
        return false
    }
    pat_len := len(pattern)
    for idx := 0; idx < len(src)-pat_len+1; idx++ {
        if src[idx:idx+pat_len] == pattern {
            return true
        }
    }
    return false
}

// MatchRegexp returns true if the string
// matches the regexp pattern, false otherwise.
func MatchRegexp(pattern string, src string) bool {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(src)
}
