package baseconstruction

import (
	"fmt"
	"math"
	"time"
)

func CalcDuration() {
	var duration string
	fmt.Println("Введите временной отрезок:")
	_, err := fmt.Scanln(&duration)
	if err != nil {
		fmt.Println(err)
		return
	}
	dateTime, err := time.ParseDuration(duration)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s = %v min\n", duration, dateTime.Minutes())
}

func LineBetweenPoint() {
	var x1, x2, y1, y2 float64
	fmt.Scan(&x1, &y1, &x2, &y2)
	d := math.Sqrt(math.Pow((x1-x2), 2) + math.Pow((y1-y2), 2))
	fmt.Println(d)
}

func RepeatString() {
	var source, result string
	var times int

	fmt.Scan(&source, &times)
	for i := 0; i < times; i++ {
		result += source
	}
	fmt.Println(result)
}

func NameByCode() {
	var code, lang string
	fmt.Scan(&code)
	switch code {
	case "en":
		lang = "English"
	case "fr":
		lang = "French"
	case "ru", "rus":
		lang = "Russian"
	default:
		lang = "Unknown"
	}
	fmt.Println(lang)
}
