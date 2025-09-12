package presentation

import (
	"fmt"

	"github.com/ilkerciblak/buldum-app/shared/core/domain"
)

// API success durumunda da bir status donebilir ama simdilik bosverelim
type ApiResult[T any] struct {
	Data           T
	ProblemDetails ProblemDetails
}

type ProblemDetails struct {
	Type   string            `json:"type"`
	Title  string            `json:"title"`
	Status int               `json:"status"`
	Detail string            `json:"detail"`
	Errors map[string]string `json:"errors,omitempty"`
}

func ToProblemDetails(e domain.IApplicationException) ProblemDetails {

	return ProblemDetails{
		Type:   fmt.Sprintf("https://httpstatuses.com/%v", e.GetCode()),
		Title:  e.GetTitle(),
		Status: e.GetCode(),
		Detail: e.GetMessage(),
		Errors: e.GetErrors(),
	}
}
