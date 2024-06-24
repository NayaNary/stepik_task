package main

import (
	"fmt"
	"sort"
	"strings"
)

// начало решения

// prettify возвращает отформатированное
// строковое представление карты
func prettify(m map[string]int) string {
	var b strings.Builder
	switch {
	case len(m) == 1:
		for key, val := range m {
			b.WriteString("{")
			b.WriteString(fmt.Sprintf(" %s: %d ", key, val))
			b.WriteString("}")
		}
		return b.String()
	case len(m) == 0:
		b.WriteString("{}")
		return b.String()
	default:
		keys := make([]string, 0, len(m))
		for key := range m {
			keys = append(keys, key)
		}
		sort.Slice(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})
		b.WriteString("{")
		b.WriteRune('\n')
		for _, key := range keys {
			b.WriteString(fmt.Sprintf("    %s: %d,", key, m[key]))
			b.WriteRune('\n')
		}
		b.WriteString("}")

		return b.String()
	}
}

// конец решения

func main() {
	m := map[string]int{}
	const want = "{}"
	got := prettify(m)
	fmt.Println(got, got == want)

}
