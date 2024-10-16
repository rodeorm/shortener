package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rodeorm/shortener/internal/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIShorten(t *testing.T) {
	type want struct {
		statusCode  int
		contentType string
	}
	tests := []struct {
		name    string
		method  string
		request string
		body    string

		server Server
		want   want
	}{
		{
			//Нужно принимать и возвращать JSON
			name:    "Проверка обработки корректных запросов: POST (json)",
			server:  Server{ServerAddress: "http://localhost:8080", Storage: repo.NewStorage("", "")},
			method:  "POST",
			body:    `{"url":"http://www.yandex.ru"}`,
			request: "http://localhost:8080/api/shorten",
			want:    want{statusCode: 201, contentType: "json"},
		},
		{
			//Нужно принимать и возвращать JSON
			name:    "Проверка обработки некорректных запросов: POST (json)",
			server:  Server{ServerAddress: "http://localhost:8080", Storage: repo.NewStorage("", "")},
			method:  "POST",
			body:    ``,
			request: "http://localhost:8080/api/shorten",
			want:    want{statusCode: 400},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var request *http.Request
			switch tt.method {
			case "POST":
				if tt.body != "" {
					fmt.Println("json", tt.body)
					request = httptest.NewRequest(http.MethodPost, tt.request, bytes.NewReader([]byte(tt.body)))

				} else {
					request = httptest.NewRequest(http.MethodPost, tt.request, nil)
				}
			case "GET":
				request = httptest.NewRequest(http.MethodGet, tt.request, nil)
			case "PUT":
				request = httptest.NewRequest(http.MethodPut, tt.request, nil)
			case "DELETE":
				request = httptest.NewRequest(http.MethodDelete, tt.request, nil)
			}
			w := httptest.NewRecorder()
			h := http.HandlerFunc(tt.server.APIShortenHandler)
			h.ServeHTTP(w, request)
			result := w.Result()
			err := result.Body.Close()
			require.NoError(t, err)
			assert.Equal(t, tt.want.statusCode, result.StatusCode)

		})
	}
}
