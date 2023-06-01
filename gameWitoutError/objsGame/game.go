package objsgame

import "fmt"

// Game описывает игру
type Game struct {
	// игрок
	player *player
	// объекты игрового мира
	things map[label]int
	// количество успешно выполненных шагов
	nSteps int
}

// newGame() создает новую игру
func NewGame() *Game {
	p := newPlayer()
	things := map[label]int{
		Apple.Name:    2,
		Coin.Name:     3,
		Mirror.Name:   1,
		Mushroom.Name: 1,
	}
	return &Game{p, things, 0}
}

// has() проверяет, остались ли в игровом мире указанные предметы
func (g *Game) has(obj Thing) bool {
	count := g.things[obj.Name]
	return count > 0
}

// Execute() выполняет шаг step
func (g *Game) Execute(st Step) error {
	// проверяем совместимость команды и объекта
	if !st.isValid() {
		return GameOverError{
			NSteps: g.nSteps,
			Err: InvalidStepError{
				Err:     fmt.Errorf("cannot %s", st.Obj.Name),
				Command: st.Cmd,
				Object:  st.Obj},
		}
	}

	// когда игрок берет или съедает предмет,
	// тот пропадает из игрового мира
	if st.Cmd == Take || st.Cmd == Eat {
		if !g.has(st.Obj) {
			return GameOverError{
				NSteps: g.nSteps,
				Err: NotEnoughObjectsError{
					Err:    fmt.Errorf("there are no %ss left", st.Obj.Name),
					Object: st.Obj},
			}
		}
		g.things[st.Obj.Name]--
	}

	// выполняем команду от имени игрока
	if err := g.player.do(st.Cmd, st.Obj); err != nil {
		return GameOverError{
			NSteps: g.nSteps,
			Err:    err,
		}
	}

	g.nSteps++
	return nil
}
