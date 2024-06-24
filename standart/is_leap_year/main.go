package main

import (
	"fmt"
	"time"
)

// начало решения

func isLeapYear(year int) bool {
	return time.Date(year,12,31,0,0,0,0, time.Local).YearDay() == 366
}

// конец решения
func main() {
	if !isLeapYear(2020) {
		fmt.Println("2020 is a leap year")
	}
	if isLeapYear(2022) {
		fmt.Println("2022 is NOT a leap year")
	}
	fmt.Println(isLeapYear(2020))
    // true

    fmt.Println(isLeapYear(2022))
    // false
	fmt.Println(isLeapYear(1900))
    // false
}
