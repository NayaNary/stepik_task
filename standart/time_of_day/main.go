package main

import (
	"errors"
	"fmt"
	"time"
)

// начало решения

// TimeOfDay описывает время в пределах одного дня
type TimeOfDay struct {
	hour int
	min  int
	sec  int
	loc  *time.Location
}

// Hour возвращает часы в пределах дня
func (t TimeOfDay) Hour() int {
	return t.hour
}

// Minute возвращает минуты в пределах часа
func (t TimeOfDay) Minute() int {
	return t.min
}

// Second возвращает секунды в пределах минуты
func (t TimeOfDay) Second() int {
	return t.sec
}

// String возвращает строковое представление времени
// в формате чч:мм:сс TZ (например, 12:34:56 UTC)
func (t TimeOfDay) String() string {
	return time.Date(0, 0, 0, t.hour, t.min, t.sec, 0, t.loc).Format("15:04:05") + " " + t.loc.String()
}

// Equal сравнивает одно время с другим.
// Если у t и other разные локации - возвращает false.
func (t TimeOfDay) Equal(other TimeOfDay) bool {
	if t.loc.String() != other.loc.String() {
		return false
	}
	tM := time.Date(0, 0, 0, t.hour, t.min, t.sec, 0, t.loc)
	fmt.Println(tM)
	otherT := time.Date(0, 0, 0, other.hour, other.min, other.sec, 0, other.loc)
	fmt.Println(otherT)
	return tM.Equal(otherT)
}

// Before возвращает true, если время t предшествует other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) Before(other TimeOfDay) (bool, error) {
	if t.loc.String() != other.loc.String() {
		return false, errors.New("разные локации")
	}
	tM := time.Date(0, 0, 0, t.hour, t.min, t.sec, 0, t.loc)
	otherT := time.Date(0, 0, 0, other.hour, other.min, other.sec, 0, other.loc)
	return tM.Before(otherT), nil
}

// After возвращает true, если время t идет после other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) After(other TimeOfDay) (bool, error) {
	if t.loc.String() != other.loc.String() {
		return false, errors.New("разные локации")
	}
	tM := time.Date(0, 0, 0, t.hour, t.min, t.sec, 0, t.loc)
	otherT := time.Date(0, 0, 0, other.hour, other.min, other.sec, 0, other.loc)
	return tM.After(otherT), nil
}

// MakeTimeOfDay создает время в пределах дня
func MakeTimeOfDay(hour, min, sec int, loc *time.Location) TimeOfDay {
	fmt.Println(loc.String())
	return TimeOfDay{
		hour: hour,
		min:  min,
		sec:  sec,
		loc:  loc,
	}
}

// конец решения
func main() {
	t1 := MakeTimeOfDay(10,11,12, time.UTC)
	t2 := MakeTimeOfDay(10,11,12, time.FixedZone("UTC", 0))
	
	// fmt.Println(t1.Hour(), t1.Minute(), t1.Second())
	// // 17 45 22
	
	// fmt.Println(t1)
	// // 17:45:22 UTC
	
	fmt.Println(t1.Equal(t2))
	// false
	
	// before, err := t1.Before(t2)
	// fmt.Println(before, err)
	// // true <nil>
	
	// after, err := t1.After(t2)
	// fmt.Println(after, err)
	// // false <nil>
}
