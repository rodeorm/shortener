package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/rodeorm/shortener/internal/core"
	"github.com/rodeorm/shortener/internal/repo"
	"github.com/rodeorm/shortener/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRootServers(t *testing.T) {
	type want struct {
		statusCode int
	}
	tests := []struct {
		name    string
		Server  Server
		method  string
		want    want
		request string
		body    string
	}{

		{
			//Эндпоинт GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
			name:    "Проверка обработки корректных GET запросов (отсутствуют данные короткого url)",
			Server:  Server{ServerAddress: "http://localhost:8080", Storage: repo.NewStorage("", "")},
			method:  "GET",
			request: "http://localhost:8080/10",
			want:    want{statusCode: 400},
		},
		{
			//Эндпоинт POST / принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле.
			name:    "Проверка обработки некорректных POST запросов",
			Server:  Server{ServerAddress: "http://localhost:8080", Storage: repo.NewStorage("", "")},
			method:  "POST",
			request: "http://localhost:8080",
			want:    want{statusCode: 400},
		},
		{
			//Эндпоинт POST / принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле.
			name:    "Проверка обработки корректных POST запросов",
			Server:  Server{ServerAddress: "http://localhost:8080", Storage: repo.NewStorage("", "")},
			method:  "POST",
			body:    "http://www.yandex.ru",
			request: "http://localhost:8080",
			want:    want{statusCode: 201},
		},
		{
			//Нужно учесть некорректные запросы и возвращать для них ответ с кодом 400 (любые кроме GET и POST)
			name:    "Проверка обработки некорректных запросов: PUT",
			Server:  Server{ServerAddress: "http://localhost:8080", Storage: repo.NewStorage("", "")},
			method:  "PUT",
			request: "http://localhost:8080",
			want:    want{statusCode: 400},
		},
		{
			//Нужно учесть некорректные запросы и возвращать для них ответ с кодом 400 (любые кроме GET и POST)
			name:    "Проверка обработки некорректных запросов: DELETE",
			Server:  Server{ServerAddress: "http://localhost:8080", Storage: repo.NewStorage("", "")},
			method:  "DELETE",
			request: "http://localhost:8080",
			want:    want{statusCode: 400},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var request *http.Request
			switch tt.method {
			case "POST":
				if tt.body != "" {
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
			h := http.HandlerFunc(tt.Server.RootHandler)
			h.ServeHTTP(w, request)
			result := w.Result()
			err := result.Body.Close()
			require.NoError(t, err)
			assert.Equal(t, tt.want.statusCode, result.StatusCode)

		})
	}
}

func TestRoot(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockStorager(ctrl)

	storage.EXPECT().InsertUser(gomock.Any()).Return(&core.User{Key: 1000}, false, nil).MaxTimes(3)
	storage.EXPECT().InsertURL("http://double.com", gomock.Any(), gomock.Any()).Return(&core.URL{Key: "short", HasBeenShorted: true}, nil)
	storage.EXPECT().InsertURL("http://err", gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("ошибка"))
	storage.EXPECT().InsertURL("http://valid.com", gomock.Any(), gomock.Any()).Return(&core.URL{Key: "short", HasBeenShorted: false}, nil)

	s := Server{Storage: storage}

	handler := http.HandlerFunc(s.RootHandler)
	srv := httptest.NewServer(handler)
	defer srv.Close()

	testCases := []struct {
		name         string
		method       string
		requestBody  string
		expectedCode int
		expectedBody string
	}{
		{name: "проверка на попытку сократить ранее сокращенный урл", method: http.MethodPost, requestBody: "http://double.com", expectedCode: http.StatusConflict},
		{name: "проверка на попытку сократить невалидный урл", method: http.MethodPost, requestBody: "http://err", expectedCode: http.StatusBadRequest},
		{name: "проверка на попытку сократить корректный урл, который не сокращали ранее", method: http.MethodPost, requestBody: "http://valid.com", expectedCode: http.StatusCreated},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := resty.New().R()
			req.Method = tc.method
			req.URL = srv.URL
			req.Body = tc.requestBody

			resp, err := req.Send()

			assert.NoError(t, err, "ошибка при попытке сделать запрос")
			assert.Equal(t, tc.expectedCode, resp.StatusCode(), "Код ответа не соответствует ожидаемому")
		})
	}

}
