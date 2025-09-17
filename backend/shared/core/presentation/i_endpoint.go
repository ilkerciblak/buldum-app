package presentation

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/shared/core/domain"
)

type IEndPoint interface {
	Path() string
	HandleRequest(w http.ResponseWriter, r *http.Request) (ApiResult[any], domain.IApplicationError)
}

// func GenerateHandlerFuncFromEndPoint(e IEndPoint) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		data, err := e.HandleRequest(w, r)
// 		if err != nil {

// 			RespondWithProblemDetails(w, err)
// 			return
// 		}

// 		if data.Data != nil {
// 			RespondWithJSON(w, data)
// 		}

// 		// TODO RESPONDWITHJSON ETC
// 	}
// }
