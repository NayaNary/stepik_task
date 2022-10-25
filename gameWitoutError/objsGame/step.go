package objsgame

// step описывает шаг игры: сочетание команды и объекта
type step struct {
    cmd command
    obj thing
}

// isValid() возвращает true, если объект
// совместим с командой
func (s step) isValid() bool {
    return s.obj.supports(s.cmd)
}