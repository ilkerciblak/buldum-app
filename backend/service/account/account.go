package account

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/api/middleware"
	"github.com/ilkerciblak/buldum-app/service/account/internal/presentation"
)

func RegisterAccountDomainEndPoints(mux *http.ServeMux) {
	mux.HandleFunc(presentation.CreateAccountEndPoint{}.Path(), middleware.ChainMiddlewaresWithEndpoint(presentation.CreateAccountEndPoint{}, middleware.LoggingMiddleware{}))
}
