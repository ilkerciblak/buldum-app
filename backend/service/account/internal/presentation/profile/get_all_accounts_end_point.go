package presentation

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/service/account/internal/application/query"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type GetAllProfilesEndPoint struct {
	Repository repository.IAccountRepository
}

func (e GetAllProfilesEndPoint) Path() string {
	return "GET /accounts"
}

func (e GetAllProfilesEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) corepresentation.ApiResult[any] {
	if r.Method != http.MethodGet {
		return *corepresentation.NewErrorResult(coredomain.MethodNotAllowed)
	}

	limit := r.URL.Query().Get("limit")
	orderBy := r.URL.Query().Get("order")
	sortBy := r.URL.Query().Get("sort")
	page := r.URL.Query().Get("page")
	isArchived := r.URL.Query().Get("is_archived")
	username := r.URL.Query().Get("user_name")

	q, err := query.NewAccountGetAllQuery(
		query.SetPage(page),
		query.SetOrderBy(orderBy),
		query.SetLimit(limit),
		query.SetSortBy(sortBy),
		query.SetIsArchived(isArchived),
		query.SetUsername(username),
	)

	if err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}

	data, handlerErr := q.Handler(e.Repository, r.Context())
	if handlerErr != nil {
		return *corepresentation.NewErrorResult(handlerErr)
	}

	return corepresentation.ApiResult[any]{
		Data:       data,
		StatusCode: http.StatusOK,
	}

}
