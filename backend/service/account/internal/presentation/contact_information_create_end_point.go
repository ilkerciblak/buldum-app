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

type ContactInformationCreateEndPoint struct {
	service application.ContactInformationServiceInterface
}

func NewContactInformationCreateEndPoint(service application.ContactInformationServiceInterface) *ContactInformationCreateEndPoint {
	return &ContactInformationCreateEndPoint{
		service: service,
	}

}

func (c ContactInformationCreateEndPoint) Path() string {
	return "POST /accounts/{user_id}/contact-informations"
}

func (c ContactInformationCreateEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) corepresentation.ApiResult[any] {

	if r.Method != http.MethodPost {
		return *corepresentation.NewErrorResult(coredomain.MethodNotAllowed)
	}

	var dto dto.ContactInformationCreateDTO

	if err := jsonmapper.DecodeRequestBody(r, dto); err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}

	userID := r.PathValue("user_id")
	parsed, err := uuid.Parse(userID)
	if err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}
	dto.SetAccountID(parsed)

	if err := c.service.Create(dto, r.Context()); err != nil {
		return *corepresentation.NewErrorResult(err)
	}

	return corepresentation.ApiResult[any]{
		Data:       nil,
		StatusCode: http.StatusCreated,
	}
}
