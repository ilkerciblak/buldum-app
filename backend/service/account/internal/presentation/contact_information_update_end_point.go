package presentation

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/dto"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
	"github.com/ilkerciblak/buldum-app/shared/helper/jsonmapper"
)

type ContactInformationUpdateEndPoint struct {
	service application.ContactInformationServiceInterface
}

func NewContactInformationUpdateEndPoint(service application.ContactInformationServiceInterface) *ContactInformationUpdateEndPoint {
	return &ContactInformationUpdateEndPoint{
		service: service,
	}
}

func (c ContactInformationUpdateEndPoint) Path() string {
	return "PUT /accounts/{user_id}/contact-informations/{id}"
}

func (c ContactInformationUpdateEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) corepresentation.ApiResult[any] {

	if r.Method != http.MethodPut {
		return *corepresentation.NewErrorResult(coredomain.MethodNotAllowed)
	}

	var dto dto.ContactInformationUpdateDTO

	if err := jsonmapper.DecodeRequestBody(r, dto); err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}

	userID := r.PathValue("user_id")
	parsed, err := uuid.Parse(userID)
	if err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}
	id := r.PathValue("id")
	parsedInstanceId, err := uuid.Parse(id)
	if err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}
	dto.SetAccountID(parsed)

	if err := c.service.Update(parsedInstanceId, dto, r.Context()); err != nil {
		return *corepresentation.NewErrorResult(err)
	}

	return corepresentation.ApiResult[any]{
		Data:       nil,
		StatusCode: http.StatusNoContent,
	}
}
