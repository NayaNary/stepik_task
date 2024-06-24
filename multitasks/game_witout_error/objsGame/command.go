package objsgame

// label - уникальное наименование
type label string

// command - команда, которую можно выполнять в игре
type command label

// список доступных команд
var (
    Eat  = command("eat")
    Take = command("take")
    Talk = command("talk to")
)