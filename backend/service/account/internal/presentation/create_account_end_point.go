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
	Repository repository.AccountRepository
}

func (c CreateAccountEndPoint) Path() string {
	return "/account"
}

func (c CreateAccountEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) (corepresentation.ApiResult[any], coredomain.IApplicationError) {

	com, err := jsonmapper.DecodeRequestBody[command.CreateAccountCommand](r)
	if err != nil {
		return corepresentation.ApiResult[any]{}, coredomain.BadRequest.WithMessage(err)
	}

	if err := com.Handler(c.Repository, r.Context()); err != nil {
		return corepresentation.ApiResult[any]{}, err
	}

	return corepresentation.ApiResult[any]{
		Data:       nil,
		StatusCode: http.StatusCreated,
	}, nil

}
