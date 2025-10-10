package presentation

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/service/account/internal/application/command"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
	"github.com/ilkerciblak/buldum-app/shared/helper/jsonmapper"
)

type UpdateAccountEndPoint struct {
	Repository repository.IAccountRepository
}

func (e UpdateAccountEndPoint) Path() string {
	return "PUT /accounts/{id}"
}
func (e UpdateAccountEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) corepresentation.ApiResult[any] {

	if r.Method != http.MethodPut {
		return *corepresentation.NewErrorResult(coredomain.MethodNotAllowed)
	}

	var cmd command.UpdateAccountCommand

	if err := jsonmapper.DecodeRequestBody(r, &cmd); err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}

	if err := cmd.SetUserID(r.PathValue("id")); err != nil {

		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}

	if err := cmd.Handler(e.Repository, r.Context()); err != nil {
		return *corepresentation.NewErrorResult(err)
	}

	return corepresentation.ApiResult[any]{
		Data:       nil,
		StatusCode: http.StatusNoContent,
	}
}
