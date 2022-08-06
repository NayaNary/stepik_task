package baseconstruction

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// Result представляет результат матча
type Result byte

// возможные результаты матча
const (
	win  Result = 'W'
	draw Result = 'D'
	loss Result = 'L'
)

var CountBall map[Result]int = map[Result]int{win: 3, draw: 1, loss: 0}
var AntonymsBall map[Result]Result = map[Result]Result{win: loss, draw: draw, loss: win}

// Team представляет команду
type Team byte

// Match представляет матч
// содержит три поля:
// - first (первая команда)
// - second (вторая команда)
// - result (результат матча)
// например, строке BAW соответствует
// first=B, second=A, result=W
type Match struct {
	First  Team
	Second Team
	Result Result
}

func (trn *Tournament) GetMatches() []Match {
	return *trn
}

// Rating представляет турнирный рейтинг команд -
// количество набранных очков по каждой команде
type Rating map[Team]int

// Tournament представляет турнир
type Tournament []Match

// CalcRating считает и возвращает рейтинг турнира
func (trn *Tournament) CalcRating() Rating {
	m := trn.GetMatches()
	result := make(Rating)
	for _, val := range m {
		amountF := result[val.First]
		amountF += CountBall[val.Result]
		result[val.First] = amountF
		amountS := result[val.Second]
		amountS+= CountBall[AntonymsBall[val.Result]]
		result[val.Second] = amountS
	}
	return result
}

// ParseTournament парсит турнир из исходной строки
func ParseTournament(s string) Tournament {
	descriptions := strings.Split(s, " ")
	trn := Tournament{}
	for _, descr := range descriptions {
		m := ParseMatch(descr)
		trn.addMatch(m)
	}
	return trn
}

// ParseMatch парсит матч из фрагмента исходной строки
func ParseMatch(s string) Match {
	return Match{
		First:  Team(s[0]),
		Second: Team(s[1]),
		Result: Result(s[2]),
	}
}

// addMatch добавляет матч к турниру
func (t *Tournament) addMatch(m Match) {
	*t = append(*t, m)
}

// Print печатает результаты турнира
// в формате Aw Bx Cy Dz
func (r *Rating) Print() {
	var parts []string
	for team, score := range *r {
		part := fmt.Sprintf("%c%d", team, score)
		parts = append(parts, part)
	}
	sort.Strings(parts)
	fmt.Println(strings.Join(parts, " "))
}

// readString считывает и возвращает строку из os.Stdin
func ReadString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return str
}
