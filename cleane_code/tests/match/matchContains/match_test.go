package matchcontains

import (
    "testing"
)

// pretty long string
const src = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

// matches in the middle of the string
const pattern = "commodo"

func BenchmarkMatchContains(b *testing.B) {
    for n := 0; n < b.N; n++ {
        MatchContains(pattern, src)
    }
}

func BenchmarkMatchContainsCustom(b *testing.B) {
    for n := 0; n < b.N; n++ {
        MatchContainsCustom(pattern, src)
    }
}

func BenchmarkMatchRegexp(b *testing.B) {
    for n := 0; n < b.N; n++ {
        MatchRegexp(pattern, src)
    }
}