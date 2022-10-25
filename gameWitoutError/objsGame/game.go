package objsgame

import "fmt"

// game описывает игру
type game struct {
	// игрок
	player *player
	// объекты игрового мира
	things map[label]int
	// количество успешно выполненных шагов
	nSteps int
}

// newGame() создает новую игру
func newGame() *game {
    p := newPlayer()
    things := map[label]int{
        apple.name:    2,
        coin.name:     3,
        mirror.name:   1,
        mushroom.name: 1,
    }
    return &game{p, things, 0}
}


// has() проверяет, остались ли в игровом мире указанные предметы
func (g *game) has(obj thing) bool {
	count := g.things[obj.name]
	return count > 0
}

// execute() выполняет шаг step
func (g *game) execute(st step) error {
	// проверяем совместимость команды и объекта
	if !st.isValid() {
		return fmt.Errorf("cannot %s", st)
	}

	// когда игрок берет или съедает предмет,
	// тот пропадает из игрового мира
	if st.cmd == take || st.cmd == eat {
		if !g.has(st.obj) {
			return fmt.Errorf("there are no %ss left", st.obj)
		}
		g.things[st.obj.name]--
	}

	// выполняем команду от имени игрока
	if err := g.player.do(st.cmd, st.obj); err != nil {
		return err
	}

	g.nSteps++
	return nil
}
