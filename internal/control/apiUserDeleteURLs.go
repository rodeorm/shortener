package control

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

/*APIUserDeleteURLsHandler принимает список идентификаторов сокращённых URL для удаления в формате: [ "a", "b", "c", "d", ...].
В случае успешного приёма запроса хендлер должен возвращать HTTP-статус 202 Accepted.
*/
func (h DecoratedHandler) APIUserDeleteURLsHandler(w http.ResponseWriter, r *http.Request) {
	w, userKey := h.GetUserIdentity(w, r)
	_, err := strconv.Atoi(userKey)
	if err != nil {
		fmt.Println("Проблемы с получением пользователя", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Проблемы с получением данных для удаления", err)
		w.WriteHeader(http.StatusNoContent)
		return
	}
	go h.Storage.DeleteURLs(string(bodyBytes), userKey)
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, string(bodyBytes))
}
