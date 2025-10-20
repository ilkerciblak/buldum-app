package presentation

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/dto"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
	"github.com/ilkerciblak/buldum-app/shared/helper/jsonmapper"
	"github.com/ilkerciblak/buldum-app/shared/logging"
)

type UpdateAccountEndPoint struct {
	Service application.AccountServiceInterface
	Logger  logging.ILogger
}

func (e UpdateAccountEndPoint) Path() string {
	return "PUT /accounts/{id}"
}

func (e UpdateAccountEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) corepresentation.ApiResult[any] {

	if r.Method != http.MethodPut {
		return *corepresentation.NewErrorResult(coredomain.MethodNotAllowed)
	}

	var dto dto.AccountUpdateDTO

	if err := jsonmapper.DecodeRequestBody(r, &dto); err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}

	if err := dto.Validate(); err != nil {
		return *corepresentation.NewErrorResult(err)
	}

	userId, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}

	if err := e.Service.UpdateAccount(userId, dto, r.Context()); err != nil {
		return *corepresentation.NewErrorResult(err)
	}

	return corepresentation.ApiResult[any]{
		Data:       nil,
		StatusCode: http.StatusNoContent,
	}
}
