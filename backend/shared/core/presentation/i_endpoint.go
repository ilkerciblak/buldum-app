package presentation

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type IEndPoint interface {
	Path() string
	HandleRequest(w http.ResponseWriter, r *http.Request) (ApiResult[any], coredomain.IApplicationError)
}
