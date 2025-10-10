package presentation

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/service/account/internal/application/command"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
	"github.com/ilkerciblak/buldum-app/shared/helper/jsonmapper"
)

type CreateAccountEndPoint struct {
	Repository repository.IAccountRepository
}

func (c CreateAccountEndPoint) Path() string {
	return "POST /accounts"
}

func (c CreateAccountEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) corepresentation.ApiResult[any] {

	if r.Method != http.MethodPost {
		return *corepresentation.NewErrorResult(coredomain.MethodNotAllowed)
	}
	var com command.CreateAccountCommand
	err := jsonmapper.DecodeRequestBody(r, &com)
	if err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}

	if err := com.Handler(c.Repository, r.Context()); err != nil {
		return *corepresentation.NewErrorResult(err)
	}

	return corepresentation.ApiResult[any]{
		Data:       nil,
		StatusCode: http.StatusCreated,
	}

}
