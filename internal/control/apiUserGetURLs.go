package control

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

/*APIUserGetURLsHandler возвращает пользователю все когда-либо сокращённые им URL в формате JSON*/
func (h DecoratedHandler) APIUserGetURLsHandler(w http.ResponseWriter, r *http.Request) {
	w, userKey := h.GetUserIdentity(w, r)
	userID, err := strconv.Atoi(userKey)
	if err != nil {
		fmt.Println("Проблемы с получением пользователя", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	URLHistory, err := h.Storage.SelectUserURLHistory(userID)
	if err != nil {
		fmt.Println("Проблемы с получением истории пользователя", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	bodyBytes, err := json.Marshal(URLHistory)
	if err != nil {
		fmt.Println("Проблемы при маршалинге истории урл", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(bodyBytes))
}
