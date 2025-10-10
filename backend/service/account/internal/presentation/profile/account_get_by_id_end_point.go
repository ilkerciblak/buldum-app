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
func (a AccountGetByIdEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) corepresentation.ApiResult[any] {
	if r.Method != http.MethodGet {
		return *corepresentation.NewErrorResult(coredomain.MethodNotAllowed)
	}

	id := r.PathValue("id")

	query, err := query.NewAccountGetByIdQuery(id)
	if err != nil {
		return *corepresentation.NewErrorResult(coredomain.NotFound.WithMessage(err))
	}

	data, err2 := query.Handler(a.Repository, r.Context())
	if err2 != nil {
		return *corepresentation.NewErrorResult(err2)
	}
	return corepresentation.ApiResult[any]{
		Data:       data,
		StatusCode: 200,
	}
}
