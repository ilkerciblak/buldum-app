package middleware

import (
	"log"
	"net/http"

	"github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type JsonResponseHandlerMiddleware struct {
}

func (e JsonResponseHandlerMiddleware) Act(endPoint presentation.IEndPoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("JsonResponseHandlerMiddleware Calisti")
		data, err := endPoint.HandleRequest(w, r)

		if err != nil {
			presentation.RespondWithProblemDetails(w, err)
			return
		}

		if data.Data != nil {
			presentation.RespondWithJSON(w, data)
		}
		w.WriteHeader(data.StatusCode)
		log.Printf("JsonResponseHandlerMiddleware Bitirdi")
	}

}
