package middleware

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestWithLog - middleware для логирования запросов и ответов
func TestWithLog(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, tested world!"))
	})

	loggedHandler := WithLog(testHandler)

	req := httptest.NewRequest(http.MethodGet, "http://example.com/test", nil)
	rr := httptest.NewRecorder()

	loggedHandler.ServeHTTP(rr, req)

	res := rr.Result()
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Ошибка при чтении тела ответа: %v", err)
	}

	assert.Equal(t, http.StatusOK, res.StatusCode, fmt.Sprintf("Код ответа не соответствует ожидаемому. Тело запроса: %s", body))
	assert.Equal(t, "Hello, tested world!", string(body), "Содержимое ответа не соответствует ожидаемому.")
}
