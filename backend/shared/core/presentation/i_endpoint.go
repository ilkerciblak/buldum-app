package corepresentation

import (
	"log"
	"net/http"

	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type IEndPoint interface {
	Path() string
	HandleRequest(w http.ResponseWriter, r *http.Request) (ApiResult[any], coredomain.IApplicationError)
}

func GenerateHandlerFuncFromEndPoint(endPoint IEndPoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("JsonResponseHandlerMiddleware Calisti")
		data, err := endPoint.HandleRequest(w, r)

		if err != nil {
			RespondWithProblemDetails(w, err)
			return
		}

		if data.Data != nil {
			RespondWithJSON(w, data.Data)
			return
		}
		w.WriteHeader(data.StatusCode)
		log.Printf("JsonResponseHandlerMiddleware Bitirdi")
	}

}
