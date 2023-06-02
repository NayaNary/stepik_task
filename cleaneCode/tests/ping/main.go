package main

import (
	"fmt"
	"net/http"
	"ping/pingV2"
)

func main() {
	client := &http.Client{}
	pinger := pingV2.Pinger{client}
	url := "https://ya.ru"
	alive := pinger.Ping(url)
	fmt.Println(url, "is alive =", alive)
}
