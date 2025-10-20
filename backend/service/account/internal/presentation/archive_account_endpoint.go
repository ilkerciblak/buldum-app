package presentation

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type ArchiveAccountEndPoint struct {
	Service application.AccountServiceInterface
}

func (e ArchiveAccountEndPoint) Path() string {
	return "POST /accounts/{id}/archive"
}

func (e ArchiveAccountEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) corepresentation.ApiResult[any] {
	if r.Method != http.MethodPost {
		return *corepresentation.NewErrorResult(coredomain.MethodNotAllowed)
	}

	id := r.PathValue("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}

	if err := e.Service.ArchiveAccount(userId, r.Context()); err != nil {
		return *corepresentation.NewErrorResult(err)
	}

	return corepresentation.ApiResult[any]{
		StatusCode: http.StatusNoContent,
	}

}
