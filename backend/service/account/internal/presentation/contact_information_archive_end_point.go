package presentation

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type ContactInformationArchiveEndPoint struct {
	service application.ContactInformationServiceInterface
}

func NewContactInformationArchiveEndPoint(service application.ContactInformationServiceInterface) *ContactInformationArchiveEndPoint {
	return &ContactInformationArchiveEndPoint{
		service: service,
	}
}

func (c ContactInformationArchiveEndPoint) Path() string {
	return "POST /contact-informations/{id}"
}

func (c *ContactInformationArchiveEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) corepresentation.ApiResult[any] {
	if r.Method != http.MethodPost {
		return *corepresentation.NewErrorResult(coredomain.MethodNotAllowed)
	}

	id := r.PathValue("id")
	parsed, err := uuid.Parse(id)
	if err != nil {
		return *corepresentation.NewErrorResult(coredomain.BadRequest.WithMessage(err))
	}

	if err := c.service.Archive(parsed, r.Context()); err != nil {
		return *corepresentation.NewErrorResult(err)
	}

	return corepresentation.ApiResult[any]{
		Data:       nil,
		StatusCode: http.StatusNoContent,
	}
}
