package corepresentation

import (
	"log"
	"net/http"
)

type IEndPoint interface {
	Path() string
	HandleRequest(w http.ResponseWriter, r *http.Request) ApiResult[any]
}

func GenerateHandlerFuncFromEndPoint(endPoint IEndPoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("JsonResponseHandlerMiddleware Calisti")
		result := endPoint.HandleRequest(w, r)

		if result.Error != nil {
			RespondWithProblemDetails(w, result.Error)
			return
		}

		if result.Data != nil {
			RespondWithJSON(w, result.Data)
			return
		}
		w.WriteHeader(result.StatusCode)
		log.Printf("JsonResponseHandlerMiddleware Bitirdi")
	}

}
