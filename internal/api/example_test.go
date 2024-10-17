package api

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/rodeorm/shortener/internal/repo"
)

func ExampleServer_apiShortenHandler() {
	server := Server{ServerAddress: "http://localhost:8080", Storage: repo.NewStorage("", "")} // С хранилищем в памяти, поэтому мокать  не надо
	body := `{"url":"http://www.yandex.ru"}`
	reqURL := "http://localhost:8080/api/shorten"

	request := httptest.NewRequest(http.MethodPost, reqURL, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	h := http.HandlerFunc(server.APIShortenHandler)
	h.ServeHTTP(w, request)
	result := w.Result()
	err := result.Body.Close()
	if err != nil {
		log.Fatal()
	}
}
