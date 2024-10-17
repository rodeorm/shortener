package api

import (
	"fmt"
	"net/http"
)

// PingDBHandler - обработчик для метода GET для маршрута /ping, который при запросе проверяет соединение с базой данных.
//
// При успешной проверке возвращает статус 200 OK, при неуспешной — 500 Internal Server Error.
func (h Server) PingDBHandler(w http.ResponseWriter, r *http.Request) {
	err := h.Storage.PingDB()
	if err == nil {
		fmt.Fprintf(w, "%s", "Успешное соединение с БД")
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", "Ошибка соединения с БД")
	}
}
