package presentation

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/query"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type AccountGetByIdEndPoint struct {
	Repository repository.AccountRepository
}

func (a AccountGetByIdEndPoint) Path() string {
	return "GET /account/{id}"
}
func (a AccountGetByIdEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) (corepresentation.ApiResult[any], coredomain.IApplicationError) {
	id := r.PathValue("id")
	userid, err := uuid.Parse(id)

	if r.Method != http.MethodGet {
		return corepresentation.ApiResult[any]{}, coredomain.MethodNotAllowed
	}

	if err != nil {
		return corepresentation.ApiResult[any]{}, coredomain.NotFound.WithMessage(err)
	}
	query := query.AccountGetByIdQuery{
		Id: userid,
	}

	data, err2 := query.Handler(a.Repository, r.Context())
	if err2 != nil {
		return corepresentation.ApiResult[any]{}, err2
	}
	return corepresentation.ApiResult[any]{Data: data, StatusCode: 200}, nil
}
