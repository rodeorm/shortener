package api

import (
	"fmt"
	"io"
	"net/http"
)

// RootHandler принимает POST запрос. В теле запроса строку надо передать URL для сокращения.
//
// Возвращает ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле, если удалось сократить URL без ошибок.
// Возвращает ответ с кодом 409, если URL был сокращен ранее.
// Возвращает ответ с кодом 400, если возникает ошибка.
func (h Server) RootHandler(w http.ResponseWriter, r *http.Request) {

	w, user, err := h.getUserIdentity(w, r)
	if err != nil {
		handleError(w, err, "RootHandler 1")
		return
	}

	bodyBytes, _ := io.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	url, err := h.Storage.InsertURL(bodyString, h.BaseURL, user)
	if err != nil {
		handleError(w, err, "RootHandler 2")
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if url.HasBeenShorted {
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	body := fmt.Sprintf("%s/%s", h.BaseURL, url.Key)
	w.Write([]byte(body))
}
