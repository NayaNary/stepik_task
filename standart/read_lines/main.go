package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// начало решения

// readLines возвращает все строки из указанного файла
func readLines(name string) ([]string, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	res := []string{}
	for _, line := range lines {
		if len(line) > 0 {
			res = append(res, strings.Trim(line, "/n"))
		}
	}
	return res, nil
}

// конец решения
func readLines2(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	res := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		res = append(res, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return res, err
	}
	return res, nil
}

func main() {
	lines, err := readLines("\\etc\\passwd")
	if err != nil {
		fmt.Println(err)
		return
	}
	for idx, line := range lines {
		fmt.Printf("%d: %s\n", idx+1, line)
	}
}
