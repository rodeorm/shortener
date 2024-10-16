package middleware

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithZip(t *testing.T) {
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	zipHandler := WithZip(testHandler)

	req := httptest.NewRequest(http.MethodPost, "http://example.com/test", io.NopCloser(strings.NewReader("Something")))
	req.Header.Add("Accept-Encoding", "gzip")
	rr := httptest.NewRecorder()

	zipHandler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Result().StatusCode, fmt.Sprintf("Код ответа не соответствует ожидаемому. Тело запроса: %s", req.Body))
	assert.Equal(t, "gzip", rr.Result().Header.Get("Content-Encoding"), fmt.Sprintf("Header не содержит Content-Encoding gzip. Header: %s", rr.Result().Header))
}
