package pingV2

import (
	"net/http"
	"strconv"
	"strings"
	"testing"
)

// MockClient - это заглушка http-клиента для тестов.
type MockClient struct{}

// Head возвращает http-ответ со статусом, указанным в url.
// Например:
// url = https://ya.ru/200 -> статус = 200
// url = https://ya.ru/404 -> статус = 404
func (client *MockClient) Head(url string) (resp *http.Response, err error) {
	parts := strings.Split(url, "/")
	last := parts[len(parts)-1]
	statusCode, err := strconv.Atoi(last)
	if err != nil {
		return nil, err
	}
	resp = &http.Response{StatusCode: statusCode}
	return resp, nil
}

func TestPing(t *testing.T) {
	client := &MockClient{}
	pinger := Pinger{client}
	got := pinger.Ping("https://example.com/200")
	if !got {
		t.Errorf("Expected example.com to be available")
	}
	got = pinger.Ping("https://example.com/404")
	if got {
		t.Errorf("Expected example.com/404 to be unavailable")
	}
}
