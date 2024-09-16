package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

/*
APIUserDeleteURLsHandler принимает список идентификаторов сокращённых URL для удаления в формате: [ "a", "b", "c", "d", ...].
В случае успешного приёма запроса хендлер должен возвращать HTTP-статус 202 Accepted.
*/
func (h Server) APIUserDeleteURLsHandler(w http.ResponseWriter, r *http.Request) {
	w, user, isUnathorized, err := h.GetUserIdentity(w, r)
	if isUnathorized {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if err != nil {
		log.Println("APIUserDeleteURLsHandler 1", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)

	if err != nil {
		log.Println("APIUserDeleteURLsHandler 2", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	go h.Storage.DeleteURLs(string(bodyBytes), user)
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, string(bodyBytes))
}
