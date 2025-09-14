package middleware

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type ErrorHandlerMiddleware struct {
}

func (e ErrorHandlerMiddleware) Act(endPoint presentation.IEndPoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := endPoint.HandleRequest(w, r)
		if err != nil {
			presentation.RespondWithProblemDetails(w, err)
			return
		}
		presentation.RespondWithJSON(w, data)
	}

}
