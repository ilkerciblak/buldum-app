package presentation

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/service/account/internal/application"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/dto"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type GetAllProfilesEndPoint struct {
	Service application.AccountServiceInterface
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

	queryDTO := dto.NewGetAllAccountDTO()
	queryDTO.SetIsArchived(isArchived)
	queryDTO.SetOrder(orderBy)
	queryDTO.SetSortBy(sortBy)
	queryDTO.SetUsername(username)
	queryDTO.SetPage(page)
	queryDTO.SetLimit(limit)

	data, err := e.Service.GetAllAccount(queryDTO.CommonQueryParameters, queryDTO.ProfileGetAllQueryFilter, r.Context())

	if err != nil {
		return *corepresentation.NewErrorResult(err)
	}

	return corepresentation.ApiResult[any]{
		Data:       data,
		StatusCode: http.StatusOK,
	}

}
