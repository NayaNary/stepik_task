package main

import (
	"fmt"
	"strconv"
	"strings"
)

// начало решения

// calcDistance возвращает общую длину маршрута в метрах
func calcDistance(directions []string) int {
	routeInMeters := 0
	for _, line := range directions {
		fields := strings.Fields(line)
		for _, field := range fields {
			if strings.Contains(field, "km") {
				num := strings.ReplaceAll(field, "km", "")
				if strings.Contains(num, ".") {
					f, _ := strconv.ParseFloat(num, 64)
					fNum := f * 1000
					routeInMeters += int(fNum)
					break
				}
				n, _ := strconv.Atoi(num)
				routeInMeters += n * 1000
				break
			}
			if strings.Contains(field, "m") {
				num := strings.ReplaceAll(field, "m", "")
				n, _ := strconv.Atoi(num)
				routeInMeters += n
				break
			}
		}
	}
	return routeInMeters
}

// конец решения

func main() {
	directions := []string{
		"straight 1.6km",
	}
	const want = 6000
	got := calcDistance(directions)
	fmt.Println(got, got == want)

}
