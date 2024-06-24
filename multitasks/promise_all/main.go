package main

import (
	"fmt"
	"time"
)

func main() {
	funcs := []func() any{squared(2), squared(3), squared(4)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}

func gather(funcs []func() any) []any {
	// выполните все переданные функции,
	// соберите результаты в срез
	// и верните его
	res := make([]any, len(funcs))
	type resFunc struct {
		index int
		val   any
	}
	ch := make(chan resFunc)
	for key, val := range funcs {
		go func(funcVal func() any, index int) {
			ch <- resFunc{index: index, val: funcVal()}
		}(val, key)
	}

	for i := 0; i < len(funcs); i++ {
		resVal:= <-ch
		res[resVal.index] = resVal.val
	}

	return res
}

func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}
