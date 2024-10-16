package api

import (
	"fmt"
	"io"
	"net/http"
)

// RootHandler POST принимает в теле запроса строку URL для сокращения и возвращает ответ с кодом 201 и сокращённым URL в виде текстовой строки в теле.
func (h Server) RootHandler(w http.ResponseWriter, r *http.Request) {

	w, user, err := h.GetUserIdentity(w, r)
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
