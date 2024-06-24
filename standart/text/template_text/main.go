package main

import (
	"bytes"
	"fmt"
	"text/template"
)

// начало решения

var templateText = `{{if ge .Balance 100 }}{{.Name}}, добрый день! Ваш баланс - {{.Balance}}₽. Все в порядке.{{else if and (gt .Balance 0) (lt .Balance 100)}}{{.Name}}, добрый день! Ваш баланс - {{.Balance}}₽. Пора пополнить.{{else if eq .Balance 0 }}{{.Name}}, добрый день! Ваш баланс - {{.Balance}}₽. Доступ заблокирован.{{end}}`

// конец решения

type User struct {
	Name    string
	Balance int
}

// renderToString рендерит данные по шаблону в строку
func renderToString(tpl *template.Template, data any) string {
	var buf bytes.Buffer
	tpl.Execute(&buf, data)
	return buf.String()
}

func main() {
	tpl := template.New("message")
	tpl = template.Must(tpl.Parse(templateText))

	user := User{"Алиса", 500}
	got := renderToString(tpl, user)

	const want = "Алиса, добрый день! Ваш баланс - 500₽. Все в порядке."
	fmt.Println(want)
	fmt.Println(got)
	fmt.Println(got == want)

}
