package presentation

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/service/account/internal/application/query"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type GetAllProfilesEndPoint struct {
	Repository repository.AccountRepository
}

func (e GetAllProfilesEndPoint) Path() string {
	return "GET /account"
}

func (e GetAllProfilesEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) (corepresentation.ApiResult[any], coredomain.IApplicationError) {
	if r.Method != http.MethodGet {
		return corepresentation.ApiResult[any]{}, coredomain.MethodNotAllowed
	}
	queryMap := corepresentation.QueryParametersMapper(r, query.AccountGetAllQuery{})
	q, err := query.NewAccountGetAllQuery(queryMap)
	if err != nil {
		return corepresentation.ApiResult[any]{}, coredomain.BadRequest.WithMessage(err)
	}

	data, handlerErr := q.Handler(e.Repository, r.Context())
	if handlerErr != nil {
		return corepresentation.ApiResult[any]{}, handlerErr
	}

	return corepresentation.ApiResult[any]{
		Data:       data,
		StatusCode: http.StatusOK,
	}, nil

}
