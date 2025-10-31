package presentation

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/dto"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type ContactInformationGetAllByUser struct {
	service application.ContactInformationServiceInterface
}

func NewContactInformationGetAllByUserEndPoint(service application.ContactInformationServiceInterface) *ContactInformationGetAllByUser {
	return &ContactInformationGetAllByUser{
		service: service,
	}
}

func (c ContactInformationGetAllByUser) Path() string {
	return "GET /accounts/{user_id}/contact-informations"
}

func (c ContactInformationGetAllByUser) HandleRequest(w http.ResponseWriter, r *http.Request) corepresentation.ApiResult[any] {
	if r.Method != http.MethodGet {
		return *corepresentation.NewErrorResult(coredomain.MethodNotAllowed)
	}

	userId := r.PathValue("user_id")
	parsed, err := uuid.Parse(userId)
	if err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}

	tp := r.URL.Query().Get("type")
	isArchived := r.URL.Query().Get("is_archived")
	isPublic := r.URL.Query().Get("is_public")

	dto := dto.NewContactInformationGetAllByAccountDTO()
	dto.SetUserID(parsed)
	dto.SetType(tp)
	dto.SetIsArchived(isArchived)
	dto.SetIsPublic(isPublic)

	data, err := c.service.GetAllFiltered(*dto, r.Context())
	if err != nil {
		return *corepresentation.NewErrorResult(err)
	}

	return corepresentation.ApiResult[any]{
		Data:       data,
		StatusCode: http.StatusOK,
	}

}
