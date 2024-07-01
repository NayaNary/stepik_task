package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// начало решения

// Duration описывает продолжительность фильма
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	timeD := time.Duration(d)
	m := int64(timeD.Minutes())
	h := int64(m / 60)
	// if m%60 > 0 {
	m = m % 60
	// }

	res := make([]byte, 0)
	rd := bytes.NewBuffer(res)
	rd.WriteString("\"")
	if h > 0 {
		// res += fmt.Sprintf("\"%dh", h)
		rd.WriteString(fmt.Sprintf("%dh", h))
	}
	if m > 0 {
		// res += fmt.Sprintf("%dm\"", m)
		rd.WriteString(fmt.Sprintf("%dm", m))
	}
	rd.WriteString("\"")

	return rd.Bytes(), nil
}

// Rating описывает рейтинг фильма
type Rating int

func (r Rating) MarshalJSON() ([]byte, error) {
	if r > 5 {
		return nil, fmt.Errorf("слишком большой рейтинг: %d", r)
	}
	res := make([]byte, 0)
	rd := bytes.NewBuffer(res)
	rd.WriteString("\"")
	rd.WriteString(strings.Repeat("★", int(r)))
	// rating := strings.Repeat("★", int(r))
	if 5-r > 0 {
		// rating += strings.Repeat("☆", int(5-r))
		rd.WriteString(strings.Repeat("☆", int(5-r)))
	}
	rd.WriteString("\"")

	return rd.Bytes(), nil
}

// Movie описывает фильм
type Movie struct {
	Title    string
	Year     int
	Director string
	Genres   []string
	Duration Duration
	Rating   Rating
}

// MarshalMovies кодирует фильмы в JSON.
//   - если indent = 0 - использует json.Marshal
//   - если indent > 0 - использует json.MarshalIndent
//     с отступом в указанное количество пробелов.
func MarshalMovies(indent int, movies ...Movie) (string, error) {
	if indent == 0 {
		res, err := json.Marshal(movies)
		return string(res), err
	}
	if indent > 0 {
		res, err := json.MarshalIndent(movies, "", strings.Repeat(" ", indent))
		return string(res), err
	}

	return "", nil
}

// конец решения

func main() {
	m1 := Movie{
		Title:    "Interstellar",
		Year:     2014,
		Director: "Christopher Nolan",
		Genres:   []string{"Adventure", "Drama", "Science Fiction"},
		Duration: Duration(2*time.Hour + 49*time.Minute),
		Rating:   5,
	}
	m2 := Movie{
		Title:    "Sully",
		Year:     2016,
		Director: "Clint Eastwood",
		Genres:   []string{"Drama", "History"},
		Duration: Duration(time.Hour + 36*time.Minute),
		Rating:   4,
	}
	m3 := Movie{
		Title:    "Sully",
		Year:     2016,
		Director: "Clint Eastwood",
		Genres:   []string{"Drama", "History"},
		Duration: Duration(time.Hour),
		Rating:   4,
	}

	s, err := MarshalMovies(4, m1, m2, m3)
	fmt.Println(err)
	// nil
	fmt.Println(s)
	/*
		[
		    {
		        "Title": "Interstellar",
		        "Year": 2014,
		        "Director": "Christopher Nolan",
		        "Genres": [
		            "Adventure",
		            "Drama",
		            "Science Fiction"
		        ],
		        "Duration": "2h49m",
		        "Rating": "★★★★★"
		    },
		    {
		        "Title": "Sully",
		        "Year": 2016,
		        "Director": "Clint Eastwood",
		        "Genres": [
		            "Drama",
		            "History"
		        ],
		        "Duration": "1h36m",
		        "Rating": "★★★★☆"
		    }
		]
	*/
}
