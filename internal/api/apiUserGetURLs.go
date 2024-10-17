package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// APIUserGetURLsHandler возвращает пользователю все когда-либо сокращённые им URL в формате JSON
func (h Server) APIUserGetURLsHandler(w http.ResponseWriter, r *http.Request) {
	w, user, err := h.getUserIdentity(w, r)

	if user.WasUnathorized {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err != nil {
		handleError(w, err, "APIUserGetURLsHandler 1")
		return
	}
	history, err := h.Storage.SelectUserURLHistory(user)
	if err != nil {
		handleError(w, err, "APIUserGetURLsHandler 1")
		return
	}

	//Не очень изящно, конечно. Т.к. не хочется слишком много мест переделывать
	for i, v := range history {
		history[i].Short = fmt.Sprintf("%s/%s", h.BaseURL, v.Short)
	}

	bodyBytes, err := json.Marshal(history)
	if err != nil {
		handleError(w, err, "APIUserGetURLsHandler 1")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(bodyBytes))
}
