package objsgame

import "fmt"

// invalidStepError - ошибка, которая возникает,
// когда команда шага не совместима с объектом
type invalidStepError string

// notEnoughObjectsError - ошибка, которая возникает,
// когда в игре закончились объекты определенного типа
type notEnoughObjectsError string

// commandLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на выполнение команды
type commandLimitExceededError string

// objectLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на количество объектов
// определенного типа в инвентаре
type objectLimitExceededError string

// gameOverError - ошибка, которая произошла в игре
type gameOverError struct {
	// количество шагов, успешно выполненных
	// до того, как произошла ошибка
	nSteps int
	err    error
	// ...
}

func (goe gameOverError) Error() string {
	return fmt.Sprintf("on strp %d error: %v", goe.nSteps, goe.err)
}

func (le gameOverError) Unwrap() error {
	return le.err
}

func giveAdvice(err error) string {
	// ...
	return ""
}
