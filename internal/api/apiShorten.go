package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rodeorm/shortener/internal/core"
)

// APIShortenHandler принимает в теле запроса JSON-объект {"url":"<some_url>"} и возвращает в ответ объект {"result":"<shorten_url>"}.
func (h Server) APIShortenHandler(w http.ResponseWriter, r *http.Request) {
	url := core.URL{}
	shortURL := core.ShortenURL{}

	w, user, err := h.GetUserIdentity(w, r)
	if err != nil {
		handleError(w, err, "APIShortenHandler 1")
		return
	}

	bodyBytes, _ := io.ReadAll(r.Body)
	err = json.Unmarshal(bodyBytes, &url)
	if err != nil {
		handleError(w, err, "APIShortenHandler 2")
		return
	}

	urlFromStorage, err := h.Storage.InsertURL(url.Key, h.BaseURL, user)
	url = *urlFromStorage
	if err != nil {
		handleError(w, err, "APIShortenHandler 3")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	shortURL.Key = h.BaseURL + "/" + url.Key
	if url.HasBeenShorted {
		w.WriteHeader(http.StatusConflict)
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	bodyBytes, err = json.Marshal(shortURL)
	if err != nil {
		handleError(w, err, "APIShortenHandler 4")
	}
	fmt.Fprint(w, string(bodyBytes))
}
