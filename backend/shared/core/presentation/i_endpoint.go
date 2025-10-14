package corepresentation

import (
	"log"
	"net/http"

	"github.com/ilkerciblak/buldum-app/shared/logging"
)

type IEndPoint interface {
	Path() string
	HandleRequest(w http.ResponseWriter, r *http.Request) ApiResult[any]
}

func GenerateHandlerFuncFromEndPoint(endPoint IEndPoint, logger logging.ILogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := endPoint.HandleRequest(w, r)

		if result.Error != nil {
			logger.With("error", ToProblemDetails(result.Error))
			RespondWithProblemDetails(w, result.Error)
			return
		}

		if result.Data != nil {
			logger.With("response", result.Data)
			RespondWithJSON(w, result.Data)
			return
		}

		w.WriteHeader(result.StatusCode)

		log.Printf("JsonResponseHandlerMiddleware Bitirdi")
	}

}
