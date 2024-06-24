package pingV2

// -- ping/ping.go --
// Пакет ping проверяет доступность URL.

import "net/http"


// HTTPClient - это усеченный http-клиент,
// который умеет делать HEAD-запросы.
type HTTPClient interface {
    Head(url string) (resp *http.Response, err error)
}

// Pinger проверяет доступность URL.
type Pinger struct {
    Client HTTPClient
}

// Ping запрашивает указанный URL.
// Возвращает true, если адрес доступен, и false в противном случае.
func (p Pinger) Ping(url string) bool {
    resp, err := p.Client.Head(url)
    if err != nil {
        return false
    }
    if resp.StatusCode != 200 {
        return false
    }
    return true
}