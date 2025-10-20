package presentation

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type AccountGetByIdEndPoint struct {
	Service application.AccountServiceInterface
}

func (a AccountGetByIdEndPoint) Path() string {
	return "GET /accounts/{id}"
}
func (a AccountGetByIdEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) corepresentation.ApiResult[any] {
	if r.Method != http.MethodGet {
		return *corepresentation.NewErrorResult(coredomain.MethodNotAllowed)
	}

	id := r.PathValue("id")

	userId, err := uuid.Parse(id)
	if err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}

	data, err := a.Service.GetAccountById(userId, r.Context())
	if err != nil {
		if k, c := err.(coredomain.IApplicationError); c {
			return *corepresentation.NewErrorResult(k)
		}

		return *corepresentation.NewErrorResult(coredomain.InternalServerError.WithMessage(err))
	}

	return corepresentation.ApiResult[any]{
		Data:       data,
		StatusCode: 200,
	}
}
