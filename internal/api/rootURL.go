package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// RootURLHandler обработчик метода GET для маршрута ./{id} принимает в качестве URL-параметра идентификатор сокращённого URL и перебрасывает по оригинальному URL.
//
// Для переброски возвращает ответ со статусом 307 TemporaryRedirect и оригинальным URL в HTTP-заголовке Location.
// Для некорректных запросов возвращает ответ со статусом 400 BadRequest.
// При запросе удалённого URL с помощью хендлера возвращается ответ со статусом 410 Gone.
func (h Server) rootURLHandler(w http.ResponseWriter, r *http.Request) {
	currentID := mux.Vars(r)["URL"]
	url, err := h.Storage.SelectOriginalURL(currentID)

	if err != nil {
		handleError(w, err, "RootHandler 1")
	}

	if url.HasBeenDeleted {
		w.WriteHeader(http.StatusGone)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Location", url.OriginalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
	fmt.Fprintf(w, "%s", url.OriginalURL)
}
