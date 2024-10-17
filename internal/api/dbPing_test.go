package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/rodeorm/shortener/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDBPing(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockStorager(ctrl)

	storage.EXPECT().PingDB().Return(nil).AnyTimes()

	s := Server{Storage: storage}

	handler := http.HandlerFunc(s.PingDBHandler)
	srv := httptest.NewServer(handler)
	defer srv.Close()

	testCases := []struct {
		name         string
		method       string
		requestURL   string
		expectedCode int
		expectedBody string
	}{
		{name: "обработка успешной попытки достучаться к БД", method: http.MethodGet, expectedCode: http.StatusOK, requestURL: "/ping"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := resty.New().R()
			req.Method = tc.method
			req.URL = srv.URL
			resp, err := req.Send()

			assert.NoError(t, err, fmt.Sprintf("ошибка при попытке сделать запрос %s, ошибка: %s", req.URL, err))
			assert.Equal(t, tc.expectedCode, resp.StatusCode(), "Код ответа не соответствует ожидаемому")
		})
	}

}
