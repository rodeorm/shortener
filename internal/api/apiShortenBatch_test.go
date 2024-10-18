package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/rodeorm/shortener/internal/core"
	"github.com/rodeorm/shortener/mocks"
)

func TestAPIShortenBatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockStorager(ctrl)

	storage.EXPECT().InsertUser(gomock.Any()).Return(&core.User{Key: 1000, WasUnathorized: false}, nil).MaxTimes(3)
	storage.EXPECT().InsertURL("http://double.com", gomock.Any(), gomock.Any()).Return(&core.URL{Key: "short", HasBeenShorted: true}, nil)
	storage.EXPECT().InsertURL("http://err", gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("ошибка"))
	storage.EXPECT().InsertURL("http://valid.com", gomock.Any(), gomock.Any()).Return(&core.URL{Key: "short", HasBeenShorted: false}, nil)

	s := Server{Storage: storage}

	handler := http.HandlerFunc(s.APIShortenBatchHandler)
	srv := httptest.NewServer(handler)
	defer srv.Close()

	testCases := []struct {
		name         string
		method       string
		requestBody  string
		expectedCode int
		expectedBody string
	}{
		{name: "проверка на попытку сократить ранее сокращенный урл", method: http.MethodPost, requestBody: "[" +
			"{" +
			"\"correlation_id\": \"<строковый идентификатор>\"," +
			"\"original_url\": \"http://double.com\"" +
			"}" +
			"]", expectedCode: http.StatusCreated}, // Непонятно, насколько правильно возвращать в этом случае "StatusCreated", т.к. другие хэндлеры для дублей возвращают конфликт
		{name: "проверка на попытку сократить невалидный урл", method: http.MethodPost, requestBody: "[" +
			"{" +
			"\"correlation_id\": \"<строковый идентификатор>\"," +
			"\"original_url\": \"http://err\"" +
			"}" +
			"]", expectedCode: http.StatusBadRequest},
		{name: "проверка на попытку сократить корректный урл, который не сокращали ранее", method: http.MethodPost, requestBody: "[" +
			"{" +
			"\"correlation_id\": \"<строковый идентификатор>\"," +
			"\"original_url\": \"http://valid.com\"" +
			"}" +
			"]", expectedCode: http.StatusCreated},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := resty.New().R()
			req.Method = tc.method
			req.URL = srv.URL
			req.Body = tc.requestBody

			resp, err := req.Send()

			assert.NoError(t, err, "ошибка при попытке сделать запрос")
			assert.Equal(t, tc.expectedCode, resp.StatusCode(), fmt.Sprintf("Код ответа не соответствует ожидаемому. Тело запроса: %s", req.Body))
		})
	}

}
