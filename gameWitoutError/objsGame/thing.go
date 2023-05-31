package objsgame

// Thing - объект, который существует в игре
type Thing struct {
	Name    label
	Actions map[command]string
}

// supports() возвращает true, если объект
// поддерживает команду action
func (t Thing) supports(action command) bool {
	_, ok := t.Actions[action]
	return ok
}

// полный список объектов в игре
var (
	Apple = Thing{"apple", map[command]string{
		Eat:  "mmm, delicious!",
		Take: "you have an apple now",
	}}
	Bob = Thing{"bob", map[command]string{
		Talk: "Bob says hello",
	}}
	Coin = Thing{"coin", map[command]string{
		Take: "you have a coin now",
	}}
	Mirror = Thing{"mirror", map[command]string{
		Take: "you have a mirror now",
		Talk: "mirror does not answer",
	}}
	Mushroom = Thing{"mushroom", map[command]string{
		Eat:  "tastes funny",
		Take: "you have a mushroom now",
	}}
)
