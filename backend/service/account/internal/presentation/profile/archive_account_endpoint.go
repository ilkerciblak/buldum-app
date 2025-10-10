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

func (e ArchiveAccountEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) corepresentation.ApiResult[any] {
	if r.Method != http.MethodPost {
		return *corepresentation.NewErrorResult(coredomain.MethodNotAllowed)
	}

	id := r.PathValue("id")

	c, err := command.NewArchiveAccountCommand(id)
	if err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}

	if err := c.Handler(e.Repository, r.Context()); err != nil {
		return *corepresentation.NewErrorResult(err)
	}

	return corepresentation.ApiResult[any]{

		StatusCode: http.StatusNoContent,
	}

}
