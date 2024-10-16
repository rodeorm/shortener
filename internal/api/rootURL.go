package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
RootURLHandler GET /{id} принимает в качестве URL-параметра идентификатор сокращённого URL и возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
Нужно учесть некорректные запросы и возвращать для них ответ с кодом 400
При запросе удалённого URL с помощью хендлера GET /{id} нужно вернуть статус 410 Gone
*/
func (h Server) RootURLHandler(w http.ResponseWriter, r *http.Request) {
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
