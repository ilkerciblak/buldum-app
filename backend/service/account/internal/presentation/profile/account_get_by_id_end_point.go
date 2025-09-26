package presentation

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/service/account/internal/application/query"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type AccountGetByIdEndPoint struct {
	Repository repository.IAccountRepository
}

func (a AccountGetByIdEndPoint) Path() string {
	return "GET /accounts/{id}"
}
func (a AccountGetByIdEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) (corepresentation.ApiResult[any], coredomain.IApplicationError) {
	if r.Method != http.MethodGet {
		return corepresentation.ApiResult[any]{}, coredomain.MethodNotAllowed
	}

	// id := r.PathValue("id")
	query, err := query.NewAccountGetByIdQuery(corepresentation.PathValuesMapper(r, query.AccountGetByIdQuery{}))
	if err != nil {
		return corepresentation.ApiResult[any]{}, coredomain.NotFound.WithMessage(err)
	}

	data, err2 := query.Handler(a.Repository, r.Context())
	if err2 != nil {
		return corepresentation.ApiResult[any]{}, err2
	}
	return corepresentation.ApiResult[any]{Data: data, StatusCode: 200}, nil
}
