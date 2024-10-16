package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/rodeorm/shortener/internal/core"
	"github.com/rodeorm/shortener/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAPIUserDeleteURLs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockStorager(ctrl)

	storage.EXPECT().InsertUser(gomock.Any()).Return(&core.User{Key: 1000}, false, nil).AnyTimes()
	storage.EXPECT().DeleteURLs(gomock.Any()).Return(nil).AnyTimes()

	s := Server{Storage: storage, DeleteQueue: &Queue{ch: make(chan *core.URL)}}

	handler := http.HandlerFunc(s.APIUserDeleteURLsHandler)
	srv := httptest.NewServer(handler)
	defer srv.Close()

	worker := NewWorker(1, s.DeleteQueue, storage, 1)
	go worker.Loop()

	testCases := []struct {
		name         string
		method       string
		requestBody  string
		expectedCode int
		expectedBody string
	}{
		{name: "проверка на попытку удалить банч урл", method: http.MethodPost, requestBody: "[" +
			"\"6qxTVvsy\", \"RTfd56hn\", \"Jlfd67ds\"" +
			"]", expectedCode: http.StatusAccepted},
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
