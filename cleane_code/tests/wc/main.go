package main

import (
	"fmt"
	"wc/data"
)

func main() {

	s := "go is awesome, php is not"
	w := data.MakeWords(s)

	fmt.Println(w.Index("go"))
	// 0
	fmt.Println(w.Index("is"))
	// 1
	fmt.Println(w.Index("is not"))
	// -1
	fmt.Println(w.Index("python"))
	// -1
}
