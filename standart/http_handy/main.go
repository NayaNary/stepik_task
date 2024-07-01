package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// начало решения

// Handy предоставляет удобный интерфейс
// для выполнения HTTP-запросов
type Handy struct {
	uri     string
	client  *http.Client
	params  url.Values
	forms   url.Values
	header  http.Header
	body    []byte
	errJSON error
}

// NewHandy создает новый экземпляр Handy
func NewHandy() *Handy {
	return &Handy{
		client: http.DefaultClient,
		params: url.Values{},
		forms:  url.Values{},
		header: http.Header{},
	}
}

// URL устанавливает URL, на который пойдет запрос
func (h *Handy) URL(uri string) *Handy {
	h.uri = uri
	return h
}

// Client устанавливает HTTP-клиента
// вместо умолчательного http.DefaultClient
func (h *Handy) Client(client *http.Client) *Handy {
	h.client = client
	return h
}

// Header устанавливает значение заголовка
func (h *Handy) Header(key, value string) *Handy {
	h.header.Add(key, value)
	return h
}

// Param устанавливает значение URL-параметра
func (h *Handy) Param(key, value string) *Handy {
	h.params.Add(key, value)
	return h
}

// Form устанавливает данные, которые будут закодированы
// как application/x-www-form-urlencoded и отправлены в теле запроса
// с соответствующим content-type
func (h *Handy) Form(form map[string]string) *Handy {
	for key, val := range form {
		h.forms.Add(key, val)
	}
	return h
}

// JSON устанавливает данные, которые будут закодированы
// как application/json и отправлены в теле запроса
// с соответствующим content-type
func (h *Handy) JSON(v any) *Handy {
	h.forms = nil
	b, err := json.Marshal(v)
	if err != nil {
		h.errJSON = err
		return h
	}
	h.header.Add("Content-Type", "application/json")
	h.header.Add("Accept", "application/json")
	h.body = b

	return h
}

// Get выполняет GET-запрос с настроенными ранее параметрами
func (h *Handy) Get() *HandyResponse {
	req, err := http.NewRequest(http.MethodGet, h.uri, nil)
	if err != nil {
		return &HandyResponse{err: err}
	}
	req.URL.RawQuery = h.params.Encode()
	req.Header = h.header
	resp, err := h.client.Do(req)
	if err != nil {
		return &HandyResponse{err: err}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	return &HandyResponse{
		StatusCode: resp.StatusCode,
		err:        err,
		body:       body,
	}
}

// Post выполняет POST-запрос с настроенными ранее параметрами
func (h *Handy) Post() *HandyResponse {
	var resp *http.Response
	var err error
	if len(h.forms) > 0 {
		resp, err = h.client.PostForm(h.uri, h.forms)
	} else {
		if h.errJSON != nil {
			return &HandyResponse{err: h.errJSON}
		}
		req, errReq := http.NewRequest(http.MethodPost, h.uri, bytes.NewReader(h.body)) // (1)
		if errReq != nil {
			return &HandyResponse{err: errReq}
		}
		req.URL.RawQuery = h.params.Encode()
		req.Header = h.header
		resp, err = h.client.Do(req)
	}
	if err != nil {
		return &HandyResponse{err: err}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return &HandyResponse{
		StatusCode: resp.StatusCode,
		err:        err,
		body:       body,
	}
}

// HandyResponse представляет ответ на HTTP-запрос
type HandyResponse struct {
	StatusCode int
	err        error
	body       []byte
}

// OK возвращает true, если во время выполнения запроса
// не произошло ошибок, а код HTTP-статуса ответа равен 200
func (r *HandyResponse) OK() bool {
	return r.err == nil && r.StatusCode == 200
}

// Bytes возвращает тело ответа как срез байт
func (r *HandyResponse) Bytes() []byte {
	return r.body
}

// String возвращает тело ответа как строку
func (r *HandyResponse) String() string {
	return string(r.body)
}

// JSON декодирует тело ответа из JSON и сохраняет
// результат по адресу, на который указывает v
func (r *HandyResponse) JSON(v any) {
	// работает аналогично json.Unmarshal()
	// если при декодировании произошла ошибка,
	// она должна быть доступна через r.Err()
	r.err = json.Unmarshal(r.body, &v)
}

// Err возвращает ошибку, которая возникла при выполнении запроса
// или обработке ответа
func (r *HandyResponse) Err() error {
	return r.err
}

// конец решения

func main() {
	//  err1 := errors.New("Новая")
	errCommon := errors.Join(nil, nil)
	fmt.Println(errCommon)

	{
		// примеры запросов

		// GET-запрос с параметрами
		NewHandy().URL("https://httpbingo.org/get").Param("id", "42").Get()

		// HTTP-заголовки
		NewHandy().
			URL("https://httpbingo.org/get").
			Header("Accept", "text/html").
			Header("Authorization", "Bearer 1234567890").
			Get()

		// POST формы
		params := map[string]string{
			"brand":    "lg",
			"category": "tv",
		}
		NewHandy().URL("https://httpbingo.org/post").Form(params).Post()

		// POST JSON-документа
		NewHandy().URL("https://httpbingo.org/post").JSON(params).Post()
	}

	{
		// пример обработки ответа

		// отправляем GET-запрос с параметрами
		resp := NewHandy().URL("https://httpbingo.org/get").Param("id", "42").Get()
		if !resp.OK() {
			panic(resp.String())
		}

		// декодируем ответ в JSON
		var data map[string]any
		resp.JSON(&data)

		fmt.Println(data["url"])
		// "https://httpbingo.org/get"
		fmt.Println(data["args"])
		// map[id:[42]]
	}
}
