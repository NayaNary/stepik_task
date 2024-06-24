package tests_test

import (
	"fmt"
	"testing"
	"tests"
)

func TestSumZero(t *testing.T) {
	if tests.Sum() != 0 {
		t.Errorf("Expected Sum() == 0")
	}
}

func TestSumOne(t *testing.T) {
	if tests.Sum(1) != 1 {
		t.Errorf("Expected Sum(1) == 1")
	}
}

func TestSumPair(t *testing.T) {
	if tests.Sum(1, 2) != 3 {
		t.Errorf("Expected Sum(1, 2) == 3")
	}
}

func TestSumMany(t *testing.T) {
	if tests.Sum(1, 2, 3, 4, 5) != 15 {
		t.Errorf("Expected Sum(1, 2, 3, 4, 5) == 15")
	}
}

func TestIntMin(t *testing.T) {
	var listTests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 1},
	}

	for _, test := range listTests {
		name := fmt.Sprintf("case(%d,%d)", test.a, test.b)
		t.Run(name, func(t *testing.T) {
			got := tests.IntMin(test.a, test.b)
			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	listTests := []struct {
		name string
		nums []int
		want int
	}{
		{"zero", []int{}, 0},
		{"one", []int{1}, 1},
		{"pair", []int{1, 2}, 3},
		{"many", []int{1, 2, 3, 4, 5}, 15},
	}
	for _, test := range listTests {
		t.Run(test.name, func(t *testing.T) {
			got := tests.Sum(test.nums...)
			if got != test.want {
				t.Errorf("%s: got %d, want %d", test.name, got, test.want)
			}
		})
	}
}

func TestSumFew(t *testing.T) {
	if tests.Sum(1, 2, 3, 4, 5) != 15 {
		t.Errorf("Expected Sum(1, 2, 3, 4, 5) == 15")
	}
}

func TestSumN(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	n := 1_000_000_000
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	got := tests.Sum(nums...)
	want := n * (n + 1) / 2
	if got != want {
		t.Errorf("Expected sum[i=1..n](i) == n*(n+1)/2")
	}
}
