package baseconstruction

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Shuffle(nums []int) {
	// перетасуйте `nums` с помощью `rand.Shuffle()``
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
}

func Filter(predicate func(int) bool, iterable []int) []int {
	// отфильтруйте `iterable` с помощью `predicate`
	// и верните отфильтрованный срез
	res := []int{}
	for _, val := range iterable {
		if predicate(val) {
			res = append(res, val)
		}
	}
	return res
}

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func ReadInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
