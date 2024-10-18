package api

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/rodeorm/shortener/internal/logger"
)

// handleError обрабатывает ошибки
func handleError(w http.ResponseWriter, err error, message string) {
	logger.Log.Error(err.Error(),
		zap.String(message, err.Error()),
	)
	w.WriteHeader(http.StatusBadRequest)
}
