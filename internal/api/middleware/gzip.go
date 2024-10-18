package middleware

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// WithLog - middleware для сжатия/распаковки
func WithZip(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		bodyBytes, _ := io.ReadAll(r.Body)
		if IsGzip(r.Header) {
			bodyBytes, _ = DecompressGzip(bodyBytes)
		}

		r.Body = io.NopCloser(strings.NewReader(string(bodyBytes)))

		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(gzipWriter{ResponseWriter: w, Writer: gz}, r)
	})
}

// gzipWriter - абстракция над Writer и ResponseWriter
type gzipWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

// Writer - синтаксическое упрощение для доступа к методу io.Writer
func (w gzipWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// DecompressGzip осуществляет декомпрессию данных, сжатых gzip
func DecompressGzip(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("ошибка при декомпрессии данных из gzip: %v", err)
	}
	defer r.Close()

	var b bytes.Buffer
	_, err = b.ReadFrom(r)
	if err != nil {
		return nil, fmt.Errorf("ошибка при декомпрессии данных из gzip: %v", err)
	}

	return b.Bytes(), nil
}

// IsGzip  проверяет по заголовкам, поддерживается ли сжатие gzip
func IsGzip(headers map[string][]string) bool {
	for _, value := range headers["Content-Encoding"] {
		if value == "application/gzip" || value == "gzip" {
			return true
		}
	}
	return false
}
