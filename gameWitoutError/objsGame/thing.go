package objsgame

// thing - объект, который существует в игре
type thing struct {
    name    label
    actions map[command]string
}

// supports() возвращает true, если объект
// поддерживает команду action
func (t thing) supports(action command) bool {
    _, ok := t.actions[action]
    return ok
}

// полный список объектов в игре
var (
    apple = thing{"apple", map[command]string{
        eat:  "mmm, delicious!",
        take: "you have an apple now",
    }}
    bob = thing{"bob", map[command]string{
        talk: "Bob says hello",
    }}
    coin = thing{"coin", map[command]string{
        take: "you have a coin now",
    }}
    mirror = thing{"mirror", map[command]string{
        take: "you have a mirror now",
        talk: "mirror does not answer",
    }}
    mushroom = thing{"mushroom", map[command]string{
        eat:  "tastes funny",
        take: "you have a mushroom now",
    }}
)