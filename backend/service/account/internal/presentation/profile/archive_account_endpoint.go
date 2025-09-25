package presentation

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/service/account/internal/application/command"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
	"github.com/ilkerciblak/buldum-app/shared/helper/jsonmapper"
)

type ArchiveAccountEndPoint struct {
	Repository repository.IAccountRepository
}

func (e ArchiveAccountEndPoint) Path() string {
	return "PUT /account"
}

func (e ArchiveAccountEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) (corepresentation.ApiResult[any], coredomain.IApplicationError) {
	// if r.Method != http.MethodPost {
	// 	return corepresentation.ApiResult[any]{}, coredomain.MethodNotAllowed
	// }

	c, err := jsonmapper.DecodeRequestBody[command.ArchiveAccountCommand](r)
	if err != nil {
		return corepresentation.ApiResult[any]{}, coredomain.BadRequest.WithMessage(err)
	}

	if err := c.Handler(e.Repository, r.Context()); err != nil {
		return corepresentation.ApiResult[any]{}, err
	}

	return corepresentation.ApiResult[any]{
		Data:       nil,
		StatusCode: http.StatusNoContent,
	}, nil

}
