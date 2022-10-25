package objsgame

import (
	"errors"
	"fmt"
)

// player - игрок
type player struct {
	// количество съеденного
	nEaten int
	// количество диалогов
	nDialogs int
	// инвентарь
	inventory []thing
}
// newPlayer создает нового игрока
func newPlayer() *player {
    return &player{inventory: []thing{}}
}

// has() возвращает true, если у игрока
// в инвентаре есть предмет obj
func (p *player) has(obj thing) bool {
	for _, got := range p.inventory {
		if got.name == obj.name {
			return true
		}
	}
	return false
}

// do() выполняет команду cmd над объектом obj
// от имени игрока
func (p *player) do(cmd command, obj thing) error {
	// действуем в соответствии с командой
	switch cmd {
	case eat:
		if p.nEaten > 1 {
			return errors.New("you don't want to eat anymore")
		}
		p.nEaten++
	case take:
		if p.has(obj) {
			return fmt.Errorf("you already have a %s", obj)
		}
		p.inventory = append(p.inventory, obj)
	case talk:
		if p.nDialogs > 0 {
			return errors.New("you don't want to talk anymore")
		}
		p.nDialogs++
	}
	return nil
}
