package middleware

import (
	"log"
	"net/http"

	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type JsonResponseHandlerMiddleware struct {
}

func (e JsonResponseHandlerMiddleware) Act(endPoint corepresentation.IEndPoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("JsonResponseHandlerMiddleware Calisti")
		data, err := endPoint.HandleRequest(w, r)

		if err != nil {
			corepresentation.RespondWithProblemDetails(w, err)
			return
		}

		if data.Data != nil {
			corepresentation.RespondWithJSON(w, data.Data)
			return
		}
		w.WriteHeader(data.StatusCode)
		log.Printf("JsonResponseHandlerMiddleware Bitirdi")
	}

}
