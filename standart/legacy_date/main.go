package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// начало решения

// asLegacyDate преобразует время в легаси-дату
func asLegacyDate(t time.Time) string {
	if t.IsZero() {
		return "0.0"
	}
	tun := t.UnixNano()
	s := tun / 1000000000
	n := tun % 1000000000
	ns := fmt.Sprintf("%d", n)
	if len(ns) < 9 {
		countZero := 9 - len(ns)
		for i := 0; i < countZero; i++ {
			if ns == "0" {
				break
			}
			ns = "0" + ns
		}
	} else {
		for i := 0; i < 9; i++ {
			if ns != "0" && strings.HasSuffix(ns, "0") {
				ns = strings.TrimSuffix(ns, "0")
			}
		}
	}

	return fmt.Sprintf("%d.%s", s, ns)
}

// parseLegacyDate преобразует легаси-дату во время.
// Возвращает ошибку, если легаси-дата некорректная.
func parseLegacyDate(d string) (time.Time, error) {
	res := strings.Split(d, ".")
	if len(res) != 2 {
		return time.Time{}, errors.New("ошибка")
	}
	s, err := strconv.Atoi(res[0])
	if err != nil {
		return time.Time{}, err
	}
	n, err := strconv.Atoi(res[1])
	if err != nil {
		return time.Time{}, err
	}
	// c := math.Pow(float64(10), float64(9-len(res[1])))
	for i := 0; i < 9-len(res[1]); i++ {
		n = n * 10
	}

	dt := time.Date(1970, 1, 1, 0, 0, s, n, time.UTC)
	return dt, nil
}

// конец решения

func main() {
	asLegacyDateSamples := map[time.Time]string{
		time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC):     "3600.123456789",
		time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC):             "3600.0",
		time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC):             "0.0",
		time.Date(2022, 5, 24, 14, 45, 22, 0, time.UTC):         "1653403522.0",
		time.Date(2022, 5, 24, 0, 0, 0, 0, time.UTC):            "1653350400.0",
		time.Date(2022, 5, 24, 14, 45, 0, 0, time.UTC):          "1653403500.0",
		time.Date(2022, 5, 24, 14, 45, 22, 951000000, time.UTC): "1653403522.951",
		time.Date(2022, 5, 24, 14, 45, 22, 951205999, time.UTC): "1653403522.951205999",
		time.Date(1970, 1, 1, 1, 0, 0, 123456000, time.UTC):     "3600.123456",
		time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC):            "1640995200.0",
		time.Date(2022, 5, 24, 14, 0, 0, 0, time.UTC):           "1653400800.0",
		time.Date(1970, 1, 1, 1, 0, 0, 1, time.UTC):             "3600.000000001",
		time.Date(1970, 1, 1, 1, 0, 0, 123456, time.UTC):        "3600.000123456",
		time.Date(2022, 5, 24, 14, 45, 22, 951205000, time.UTC): "1653403522.951205",
		time.Date(1970, 1, 1, 1, 0, 0, 123, time.UTC):     "3600.000000123",
	}
	for src, want := range asLegacyDateSamples {
		got := asLegacyDate(src)
		fmt.Println(got, want, got == want)
	}
	// parseLegacyDateSamples := map[string]time.Time{
	// 	// "3600.123456789":    time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC),
	// 	// "3600.0":            time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC),
	// 	// "0.0":               time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	// "1.123456789":       time.Date(1970, 1, 1, 0, 0, 1, 123456789, time.UTC),
	// 	// "1653403522.951000": time.Date(2022, 5, 24, 14, 45, 22, 951, time.UTC),
	// 	"3600.123":          time.Date(1970, 1, 1, 1, 0, 0, 123000000, time.UTC),
	// 	"3600.123456":       time.Date(1970, 1, 1, 1, 0, 0, 123456000, time.UTC),
	// 	"1653403522.951205": time.Date(2022, 5, 24, 14, 45, 22, 951205000, time.UTC),
	// 	"1653403522.951":    time.Date(2022, 5, 24, 14, 45, 22, 951000000, time.UTC),
	// }
	// for src, want := range parseLegacyDateSamples {
	// 	got, err := parseLegacyDate(src)
	// 	if err != nil {
	// 		fmt.Printf("%v: unexpected error\n", src)
	// 		continue
	// 	}
	// 	fmt.Println(got, want, got.Equal(want))
	// }
}
