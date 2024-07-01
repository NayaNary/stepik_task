package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// начало решения

// Email описывает письмо
type Email struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
}

// FilterEmails читает все письма из src и записывает в dst тех,
// кто проходит проверку predicate
func FilterEmails(dst io.Writer, src io.Reader, predicate func(e Email) bool) (int, error) {
	total := 0
	dec := json.NewDecoder(src)

	for {
		// декодируем очередной объект
		var email Email
		err := dec.Decode(&email)
		if err == io.EOF {
			break
		}
		if err != nil {
			return total, err
		}
		// fmt.Println(email)
		if !predicate(email) {
			continue
		}

		enc := json.NewEncoder(dst)
		err = enc.Encode(email)
		if err != nil {
			return total, err
		}
		total++
	}
	return total, nil
}

// конец решения

func main() {
	src := strings.NewReader(`
		{ "from": "alice@go.dev",      "to": "zet@php.net",              "subject": "How are you?" }
		{ "from": "bob@temp-mail.org", "to": "yolanda@java.com",         "subject": "Re: Indonesia" }
		{ "from": "cindy@go.dev",      "to": "xavier@rust-lang.org",     "subject": "Go vs Rust" }
		{ "from": "diane@dart.dev",    "to": "wanda@typescriptlang.org", "subject": "Our crypto startup" }
	`)
	dst := os.Stdout

	predicate := func(email Email) bool {
		return !strings.Contains(email.Subject, "crypto")
	}

	n, err := FilterEmails(dst, src, predicate)
	if err != nil {
		panic(err)
	}
	fmt.Println(n, "good emails")

	// {"from":"alice@go.dev","to":"zet@php.net","subject":"How are you?"}
	// {"from":"bob@temp-mail.org","to":"yolanda@java.com","subject":"Re: Indonesia"}
	// {"from":"cindy@go.dev","to":"xavier@rust-lang.org","subject":"Go vs Rust"}
	// 3 good emails
}
