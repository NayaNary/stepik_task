package objsgame

// label - уникальное наименование
type label string

// command - команда, которую можно выполнять в игре
type command label

// список доступных команд
var (
    eat  = command("eat")
    take = command("take")
    talk = command("talk to")
)