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
	inventory []Thing
}

// newPlayer создает нового игрока
func newPlayer() *player {
	return &player{inventory: []Thing{}}
}

// has() возвращает true, если у игрока
// в инвентаре есть предмет obj
func (p *player) has(obj Thing) bool {
	for _, got := range p.inventory {
		if got.Name == obj.Name {
			return true
		}
	}
	return false
}

// do() выполняет команду cmd над объектом obj
// от имени игрока
func (p *player) do(cmd command, obj Thing) error {
	// действуем в соответствии с командой
	switch cmd {
	case Eat:
		if p.nEaten > 1 {
			return InvalidStepError{
				Err:     errors.New("you don't want to eat anymore"),
				Command: cmd,
				Object:  obj}
		}
		p.nEaten++
	case Take:
		if p.has(obj) {
			return InvalidStepError{
				Err:     fmt.Errorf("you already have a %s", obj),
				Command: cmd,
				Object:  obj}
		}
		p.inventory = append(p.inventory, obj)
	case Talk:
		if p.nDialogs > 0 {
			return InvalidStepError{
				Err:     errors.New("you don't want to talk anymore"),
				Command: cmd,
				Object:  obj}
		}
		p.nDialogs++
	}
	return nil
}
