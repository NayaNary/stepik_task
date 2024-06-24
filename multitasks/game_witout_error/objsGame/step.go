package objsgame

// step описывает шаг игры: сочетание команды и объекта
type Step struct {
	Cmd command
	Obj Thing
}

// isValid() возвращает true, если объект
// совместим с командой
func (s Step) isValid() bool {
	return s.Obj.supports(s.Cmd)
}
