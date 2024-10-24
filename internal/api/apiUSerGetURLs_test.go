package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/rodeorm/shortener/internal/core"
	"github.com/rodeorm/shortener/mocks"
	"github.com/stretchr/testify/assert"
)

func TestAPIUserGetURLs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storage := mocks.NewMockStorager(ctrl)

	userURLs := make([]core.UserURLPair, 0)
	userURLs = append(userURLs, core.UserURLPair{UserKey: 1000, Short: "1", Origin: "http://1.ru"})
	userURLs = append(userURLs, core.UserURLPair{UserKey: 1000, Short: "2", Origin: "http://2.com"})

	user := &core.User{Key: 1000, WasUnathorized: false}

	storage.EXPECT().InsertUser(gomock.Any()).Return(user, nil).AnyTimes()
	storage.EXPECT().SelectUserURLHistory(user).Return(userURLs, nil)

	s := Server{Storage: storage, BaseURL: "http:tiny.com"}

	handler := http.HandlerFunc(s.APIUserGetURLsHandler)
	srv := httptest.NewServer(handler)
	defer srv.Close()

	testCases := []struct {
		name         string
		method       string
		expectedCode int
		expectedBody string
	}{
		{name: "проверка на попытку получить историю пользователя", method: http.MethodGet, expectedCode: http.StatusAccepted},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			req := resty.New().R()
			req.Method = tc.method
			req.URL = srv.URL

			resp, err := req.Send()
			assert.JSONEq(t, `[ {"short_url": "http:tiny.com/1", "original_url": "http://1.ru" }, {"short_url": "http:tiny.com/2" , "original_url": "http://2.com"}]`, string(resp.Body()))
			assert.NoError(t, err, "ошибка при попытке сделать запрос", resp)
		})
	}

}
