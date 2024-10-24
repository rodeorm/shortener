package middleware

import (
	"net/http"
	"time"

	"github.com/rodeorm/shortener/internal/logger"
	"go.uber.org/zap"
)

// WithLog - middleware для логирования запросов и ответов
func WithLog(h http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		responseData := &responseData{
			status: 0,
			size:   0,
		}
		lw := loggingResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}
		h.ServeHTTP(&lw, r)

		duration := time.Since(start)

		logger.Log.Info("request",
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
			zap.Float64("duration_ms", duration.Seconds()*1000),
			zap.Int("status", responseData.status),
			zap.Int("size", responseData.size),
		)
	}
	return http.HandlerFunc(logFn)
}

type (
	//responseData - данные ответа для логирования
	responseData struct {
		status int
		size   int
	}
	//loggingResponseWriter - декорированная абстракция для ResponseWriter
	loggingResponseWriter struct {
		http.ResponseWriter
		responseData *responseData
	}
)

// Write записывает размер
func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size // захватываем размер
	return size, err
}

// WriteHeader записывает код статуса
func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode // захватываем код статуса
}
