package main

import (
	"fmt"
	objsgame "gameWithoutError/objsGame"
	"os"
)

func main() {
	gm := objsgame.NewGame()
	// steps := []objsgame.Step{
	// 	{objsgame.Eat, objsgame.Apple},
	// 	{objsgame.Talk, objsgame.Bob},
	// 	{objsgame.Take, objsgame.Coin},
	// 	{obj
	// steps := []objsgame.Step{
	// 	{objsgame.Talk, objsgame.Bob},
	// 	{objsgame.Talk, objsgame.Bob},
	// }

	steps := []objsgame.Step{
		{objsgame.Take, objsgame.Mirror},
		{objsgame.Take, objsgame.Mirror},
	}

	for _, st := range steps {
		if err := tryStep(gm, st); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("You win!")
}

// tryStep() выполняет шаг игры и печатает результат
func tryStep(gm *objsgame.Game, st objsgame.Step) error {
	err := gm.Execute(st)
	if err != nil {
		fmt.Printf("trying to %s %s ...FAIL\n", st.Cmd, st.Obj.Name)
		fmt.Println(objsgame.GiveAdvice(err))
		return err
	}
	fmt.Printf("trying to %s %s ...OK\n", st.Cmd, st.Obj.Name)
	return err
}
