package main

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	// var text string

	// n, err := fmt.Scanln(&text)
	// fmt.Println(n, err)

	// fmt.Println(text)
	// text:="go is awesome"
	// reader := strings.NewReader(text)
	scanner := bufio.NewScanner(os.Stdin)
	// lines := []string{}
	// scanner.Split(bufio.ScanWords)
	c := cases.Title(language.Und, cases.Compact)
	// for scanner.Scan() {
	// lines = append(lines, c.String(scanner.Text()))
	// }
	// c.String(scanner.Text())
	// if err := scanner.Err(); err != nil {
	// 	panic(err)
	// }
	scanner.Scan()
	fmt.Println(c.String(scanner.Text()))
}
