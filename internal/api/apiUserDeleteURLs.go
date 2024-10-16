package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/rodeorm/shortener/internal/core"
)

/*
Хендлер DELETE /api/user/urls, который в теле запроса принимает список идентификаторов сокращённых URL для асинхронного удаления. Запрос может быть таким:
DELETE http://localhost:8080/api/user/urls
Content-Type: application/json

["6qxTVvsy", "RTfd56hn", "Jlfd67ds"]
В случае успешного приёма запроса хендлер должен возвращать HTTP-статус 202 Accepted. Фактический результат удаления может происходить позже — оповещать пользователя об успешности или неуспешности не нужно.
Успешно удалить URL может пользователь, его создавший.
*/
func (h Server) APIUserDeleteURLsHandler(w http.ResponseWriter, r *http.Request) {
	w, user, err := h.GetUserIdentity(w, r)
	if user.WasUnathorized {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		handleError(w, err, "APIUserDeleteURLsHandler 1")
		return
	}

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		handleError(w, err, "APIUserDeleteURLsHandler 2")
		return
	}

	// Помещаем URL в очередь на асинхронное удаление. В случае успешного приёма запроса хендлер должен возвращать HTTP-статус 202 Accepted.
	urls, err := core.GetURLsFromString(string(bodyBytes), user)

	if err != nil {
		handleError(w, err, "APIUserDeleteURLsHandler 3")
		return
	}
	err = h.DeleteQueue.Push(urls)
	if err != nil {
		handleError(w, err, "APIUserDeleteURLsHandler 4")
		return
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, string(bodyBytes))
}
