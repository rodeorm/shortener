package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rodeorm/shortener/internal/core"
)

/*
	Хендлер POST /api/shorten/batch, принимающий в теле запроса множество URL для сокращения в формате:

[

	{
	    "correlation_id": "<строковый идентификатор>",
	    "original_url": "<URL для сокращения>"
	},
	...

]
В качестве ответа хендлер должен возвращать данные в формате:
[

	{
	    "correlation_id": "<строковый идентификатор из объекта запроса>",
	    "short_url": "<результирующий сокращённый URL>"
	},
	...

]
*/
func (h Server) APIShortenBatchHandler(w http.ResponseWriter, r *http.Request) {
	w, user, err := h.GetUserIdentity(w, r)

	if err != nil {
		handleError(w, err, "APIShortenBatch 1")
		return
	}

	var (
		urlReq []core.URLWithCorrelationRequest
		urlRes []core.URLWithCorrelationResponse
	)

	bodyBytes, _ := io.ReadAll(r.Body)
	err = json.Unmarshal(bodyBytes, &urlReq)
	if err != nil {
		handleError(w, err, "APIShortenBatch 2")
		return
	}

	for _, value := range urlReq {
		url, err := h.Storage.InsertURL(value.Origin, h.BaseURL, user)
		if err != nil {
			handleError(w, err, "APIShortenBatch 3")
			return
		}
		urlResPart := core.URLWithCorrelationResponse{CorID: value.CorID, Short: h.BaseURL + "/" + url.Key}
		urlRes = append(urlRes, urlResPart)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	bodyBytes, err = json.Marshal(urlRes)
	if err != nil {
		handleError(w, err, "APIShortenBatch 4")
	}
	fmt.Fprint(w, string(bodyBytes))
}
