package presentation

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/service/account/internal/application"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/dto"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
	"github.com/ilkerciblak/buldum-app/shared/helper/jsonmapper"
	"github.com/ilkerciblak/buldum-app/shared/logging"
)

type CreateAccountEndPoint struct {
	Service application.AccountServiceInterface
	Logger  logging.ILogger
}

func (c CreateAccountEndPoint) Path() string {
	return "POST /accounts"
}

func (c CreateAccountEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) corepresentation.ApiResult[any] {

	if r.Method != http.MethodPost {
		return *corepresentation.NewErrorResult(coredomain.MethodNotAllowed)
	}

	var createDTO dto.AccountCreateDTO
	err := jsonmapper.DecodeRequestBody(r, &createDTO)
	if err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}

	if err := createDTO.Validate(); err != nil {
		return *corepresentation.NewErrorResult(err)
	}

	if err := c.Service.CreateAccount(createDTO, r.Context()); err != nil {
		return *corepresentation.NewErrorResult(err)
	}

	return corepresentation.ApiResult[any]{
		Data:       nil,
		StatusCode: http.StatusCreated,
	}

}
