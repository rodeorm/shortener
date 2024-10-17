package api

import (
	"net/http"
)

// BadRequestHandler обрабатывает некорректные запросы и возвращает статус 400 BadRequest
func (h Server) badRequestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}
