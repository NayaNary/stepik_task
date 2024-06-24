package weatherservice

import (
	"fmt"
	"testing"
	"time"
)

type MockWeatherService struct {
	deg int
}

func (w *MockWeatherService) Forecast() int {
	return w.deg
}

type testCase struct {
	deg  int
	want string
}

var tests []testCase = []testCase{
	{-10, "холодно"},
	{0, "холодно"},
	{5, "холодно"},
	{10, "прохладно"},
	{15, "идеально"},
	{20, "жарко"},
}

func TestForecast(t *testing.T) {
	for _, test := range tests {
		name := fmt.Sprintf("%v", test.deg)
		t.Run(name, func(t *testing.T) {
			weather := Weather{service: &MockWeatherService{deg: test.deg}}
			got := weather.Forecast()
			if got != test.want {
				t.Errorf("%s: got %s, want %s", name, got, test.want)
			}
		})
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Setup tests...")
	start := time.Now()

	m.Run()

	fmt.Println("Teardown tests...")
	end := time.Now()
	fmt.Println("Tests took", end.Sub(start))

}
