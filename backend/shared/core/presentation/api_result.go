package corepresentation

import (
	"fmt"
	"net/http"

	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

// API success durumunda da bir status donebilir ama simdilik bosverelim
type ApiResult[T any] struct {
	Data       T
	StatusCode int
	Error      error
}

func NewApiResult[T any](data T, statusCode int) ApiResult[T] {
	return ApiResult[T]{
		Data:       data,
		StatusCode: statusCode,
	}
}

func NewErrorResult(err error) *ApiResult[any] {

	if appError, k := err.(coredomain.IApplicationError); k {
		return &ApiResult[any]{
			Data:       nil,
			StatusCode: appError.GetCode(),
			Error:      appError,
		}
	}

	return &ApiResult[any]{
		StatusCode: http.StatusInternalServerError,
		Error:      coredomain.InternalServerError.WithMessage(err),
	}

}

type ProblemDetails struct {
	Type   string            `json:"type"`
	Title  string            `json:"title"`
	Status int               `json:"status"`
	Detail string            `json:"detail"`
	Errors map[string]string `json:"errors,omitempty"`
}

func ToProblemDetails(err error) ProblemDetails {
	switch e := err.(type) {
	case coredomain.IApplicationError:

		return ProblemDetails{
			Type:   fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/%v", e.GetCode()),
			Title:  e.GetTitle(),
			Status: e.GetCode(),
			Detail: e.GetMessage(),
			Errors: e.GetErrors(),
		}
	default:
		t := coredomain.InternalServerError
		return ProblemDetails{
			Type:   fmt.Sprintf("https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/%v", t.GetCode()),
			Title:  t.GetTitle(),
			Status: t.GetCode(),
			Detail: t.GetMessage(),
			Errors: t.GetErrors(),
		}
	}

}
