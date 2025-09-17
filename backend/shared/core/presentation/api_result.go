package presentation

import (
	"fmt"

	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

// API success durumunda da bir status donebilir ama simdilik bosverelim
type ApiResult[T any] struct {
	Data       T
	StatusCode int
}

func NewApiResult[T any](data T, statusCode int) ApiResult[T] {
	return ApiResult[T]{
		Data:       data,
		StatusCode: statusCode,
	}
}

type ProblemDetails struct {
	Type   string            `json:"type"`
	Title  string            `json:"title"`
	Status int               `json:"status"`
	Detail string            `json:"detail"`
	Errors map[string]string `json:"errors,omitempty"`
}

func ToProblemDetails(e coredomain.IApplicationError) ProblemDetails {
	return ProblemDetails{
		Type:   fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/%v", e.GetCode()),
		Title:  e.GetTitle(),
		Status: e.GetCode(),
		Detail: e.GetMessage(),
		Errors: e.GetErrors(),
	}
}
