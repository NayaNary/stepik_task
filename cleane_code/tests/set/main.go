package main

import (
	"fmt"
	"set/data"
)

func main() {
	set := data.MakeIntSet()

	set.Add(5)
	fmt.Println(set.Contains(5))
	// true

	fmt.Println(set.Contains(42))
	// false

	// элементы множества уникальны,
	// добавить дважды один и тот же элемент не получится
	added := set.Add(5)
	fmt.Println(added)
	// false
}
