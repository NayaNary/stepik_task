package objsgame

import (
	"errors"
	"fmt"
)

// InvalidStepError - ошибка, которая возникает,
// когда команда шага не совместима с объектом
type InvalidStepError struct {
	Err     error
	Command command
	Object  Thing
}

func (ilse InvalidStepError) Error() string {
	return ilse.Err.Error()
}

func (ilse InvalidStepError) Unwrap() error {
	return ilse.Err
}

// NotEnoughObjectsError - ошибка, которая возникает,
// когда в игре закончились объекты определенного типа
type NotEnoughObjectsError struct {
	Err    error
	Object Thing
}

func (neoe NotEnoughObjectsError) Error() string {
	return neoe.Err.Error()
}

func (ilse NotEnoughObjectsError) Unwrap() error {
	return ilse.Err
}

// CommandLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на выполнение команды
type CommandLimitExceededError struct {
	Err error
}

func (clee CommandLimitExceededError) Error() string {
	return clee.Err.Error()
}

func (ilse CommandLimitExceededError) Unwrap() error {
	return ilse.Err
}

// ObjectLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на количество объектов
// определенного типа в инвентаре
type ObjectLimitExceededError struct {
	Err error
}

func (olee ObjectLimitExceededError) Error() string {
	return olee.Err.Error()
}

func (ilse ObjectLimitExceededError) Unwrap() error {
	return ilse.Err
}

// GameOverError - ошибка, которая произошла в игре
type GameOverError struct {
	// количество шагов, успешно выполненных
	// до того, как произошла ошибка
	NSteps int
	Err    error
}

func (goe GameOverError) Error() string {
	return goe.Err.Error()
}

func (goe GameOverError) Unwrap() error {
	return goe.Err
}

func GiveAdvice(err error) string {
	var ise InvalidStepError
	if errors.As(err, &ise) {
		return fmt.Sprintf("things like '%s %s' are impossible", ise.Command, ise.Object.Name)
	}
	var neoe NotEnoughObjectsError
	if errors.As(err, &neoe) {
		return fmt.Sprintf("be careful with scarce %s", neoe.Object.Name)
	}
	return ""
}
