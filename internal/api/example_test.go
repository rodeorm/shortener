package api

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/go-resty/resty/v2"
	"github.com/golang/mock/gomock"
	"github.com/rodeorm/shortener/internal/core"
	"github.com/rodeorm/shortener/internal/repo"
	"github.com/rodeorm/shortener/mocks"
)

func ExampleServer_APIShortenHandler() {
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

func ExampleServer_APIShortenBatchHandler() {
	ctrl := gomock.NewController(nil)
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

	useCases := []struct {
		name         string
		method       string
		requestBody  string
		expectedCode int
		expectedBody string
	}{
		{name: "попытка сократить ранее сокращенный урл", method: http.MethodPost, requestBody: "[" +
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

	for _, uc := range useCases {
		req := resty.New().R()
		req.Method = uc.method
		req.URL = srv.URL
		req.Body = uc.requestBody

		resp, err := req.Send()
		if err == nil {
			fmt.Println(resp)
		}
	}
}

func ExampleServer_APIUserDeleteURLsHandler() {
	ctrl := gomock.NewController(nil)
	defer ctrl.Finish()

	storage := mocks.NewMockStorager(ctrl)

	storage.EXPECT().InsertUser(gomock.Any()).Return(&core.User{Key: 1000, WasUnathorized: false}, nil).AnyTimes()
	storage.EXPECT().DeleteURLs(gomock.Any()).Return(nil).AnyTimes()

	s := Server{Storage: storage, DeleteQueue: &Queue{ch: make(chan *core.URL)}}

	handler := http.HandlerFunc(s.APIUserDeleteURLsHandler)
	srv := httptest.NewServer(handler)
	defer srv.Close()

	worker := NewWorker(1, s.DeleteQueue, storage, 1)
	go worker.loop()

	useCases := []struct {
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

	for _, tc := range useCases {

		req := resty.New().R()
		req.Method = tc.method
		req.URL = srv.URL
		req.Body = tc.requestBody

		resp, err := req.Send()
		if err == nil {
			fmt.Println(resp)
		}
	}
}

func ExampleServer_APIUserGetURLsHandler() {
	ctrl := gomock.NewController(nil)
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

	useCases := []struct {
		name         string
		method       string
		expectedCode int
		expectedBody string
	}{
		{name: "проверка на попытку получить историю пользователя", method: http.MethodGet, expectedCode: http.StatusAccepted},
	}

	for _, tc := range useCases {
		req := resty.New().R()
		req.Method = tc.method
		req.URL = srv.URL

		resp, err := req.Send()
		if err == nil {
			fmt.Println(resp)
		}
	}
}

func ExampleServer_PingDBHandler() {
	ctrl := gomock.NewController(nil)
	defer ctrl.Finish()

	storage := mocks.NewMockStorager(ctrl)

	storage.EXPECT().Ping().Return(nil).AnyTimes()

	s := Server{Storage: storage}

	handler := http.HandlerFunc(s.PingDBHandler)
	srv := httptest.NewServer(handler)
	defer srv.Close()

	useCases := []struct {
		name         string
		method       string
		requestURL   string
		expectedCode int
		expectedBody string
	}{
		{name: "обработка успешной попытки достучаться к БД", method: http.MethodGet, expectedCode: http.StatusOK, requestURL: "/ping"},
	}

	for _, tc := range useCases {
		req := resty.New().R()
		req.Method = tc.method
		req.URL = srv.URL
		resp, err := req.Send()
		if err == nil {
			fmt.Println(resp)
		}
	}
}

func ExampleServer_RootHandler() {
	ctrl := gomock.NewController(nil)
	defer ctrl.Finish()

	storage := mocks.NewMockStorager(ctrl)

	storage.EXPECT().InsertUser(gomock.Any()).Return(&core.User{Key: 1000, WasUnathorized: false}, nil).MaxTimes(3)
	storage.EXPECT().InsertURL("http://double.com", gomock.Any(), gomock.Any()).Return(&core.URL{Key: "short", HasBeenShorted: true}, nil)
	storage.EXPECT().InsertURL("http://err", gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("ошибка"))
	storage.EXPECT().InsertURL("http://valid.com", gomock.Any(), gomock.Any()).Return(&core.URL{Key: "short", HasBeenShorted: false}, nil)

	s := Server{Storage: storage}

	handler := http.HandlerFunc(s.RootHandler)
	srv := httptest.NewServer(handler)
	defer srv.Close()

	useCases := []struct {
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

	for _, tc := range useCases {
		req := resty.New().R()
		req.Method = tc.method
		req.URL = srv.URL
		req.Body = tc.requestBody

		resp, err := req.Send()
		if err == nil {
			fmt.Println(resp)
		}
	}
}

func ExampleServer_RootURLHandler() {
	ctrl := gomock.NewController(nil)
	defer ctrl.Finish()

	storage := mocks.NewMockStorager(ctrl)

	storage.EXPECT().SelectOriginalURL(gomock.Any()).Return(&core.URL{Key: "Short", HasBeenDeleted: false}, nil).AnyTimes()

	s := Server{Storage: storage}

	handler := http.HandlerFunc(s.RootURLHandler)
	srv := httptest.NewServer(handler)
	defer srv.Close()

	useCases := []struct {
		name         string
		method       string
		requestURL   string
		expectedCode int
		expectedBody string
	}{
		{name: "редирект, если URL был сокращен ранее", method: http.MethodGet, expectedCode: http.StatusTemporaryRedirect, requestURL: "http://www.yandex.ru"},
	}

	for _, tc := range useCases {
		req := resty.New().R()
		req.Method = tc.method
		req.URL = srv.URL + "/" + tc.requestURL
		resp, err := req.Send()
		if err == nil {
			fmt.Println(resp)
		}
	}
}
