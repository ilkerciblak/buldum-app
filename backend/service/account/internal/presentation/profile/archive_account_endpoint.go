package presentation

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/service/account/internal/application/command"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type ArchiveAccountEndPoint struct {
	Repository repository.IAccountRepository
}

func (e ArchiveAccountEndPoint) Path() string {
	return "POST /accounts/{id}/archive"
}

func (e ArchiveAccountEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) (corepresentation.ApiResult[any], coredomain.IApplicationError) {
	if r.Method != http.MethodPost {
		return corepresentation.ApiResult[any]{}, coredomain.MethodNotAllowed
	}

	m := corepresentation.PathValuesMapper(r, command.ArchiveAccountCommand{})
	c, err := command.NewArchiveAccountCommand(m)
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
