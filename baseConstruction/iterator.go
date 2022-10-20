package baseconstruction

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"stepik_task/baseConstruction/interfaces"
	"strconv"
)

// weightFunc - функция, которая возвращает вес элемента
type WeightFunc func(interfaces.Element) int

// intIterator - итератор по целым числам
// (реализует интерфейс iterator)
type intIterator struct {
	// поля структуры
	val   interfaces.Element
	len   int
	index int
}

// методы intIterator, которые реализуют интерфейс iterator

// конструктор intIterator
func newIntIterator(src []int) *intIterator {
	return &intIterator{val: src, len: len(src), index: -1}
}

func (it *intIterator) Next() bool {
	it.index = it.index + 1
	return it.index < it.len
}

func (it *intIterator) Val() interfaces.Element {
	vals := it.val.([]int)
	return vals[it.index]
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// main находит максимальное число из переданных на вход программы.
func StartGetMaxElement() {
	// nums := readInput()
	nums := []int{1, 4, 5, 2, 3}
	it := newIntIterator(nums)
	weight := func(el interfaces.Element) int {
		return el.(int)
	}
	m := max(it, weight)
	fmt.Println(m)
}

// max возвращает максимальный элемент в последовательности.
// Для сравнения элементов используется вес, который возвращает
// функция weight.
func max(it interfaces.Iterator, weight WeightFunc) interfaces.Element {
	var maxEl interfaces.Element
	for it.Next() {
		curr := it.Val()
		fmt.Println(curr)
		if maxEl == nil || weight(curr) > weight(maxEl) {
			maxEl = curr
		}
	}
	return maxEl
}

// readInput считывает последовательность целых чисел из os.Stdin.
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	return nums
}
