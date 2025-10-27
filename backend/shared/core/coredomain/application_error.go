package coredomain

import (
	"fmt"
	"net/http"
)

type IApplicationError interface {
	GetTitle() string
	GetCode() int
	GetErrors() map[string]string
	GetMessage() string
	error
}

type ApplicationError struct {
	// Title
	Code    int               `json:"code"`
	Title   string            `json:"title,omitempty"`
	Message string            `json:"message,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}

func (a ApplicationError) GetTitle() string {
	return a.Title
}

func (a ApplicationError) GetCode() int {
	return a.Code
}

func (a ApplicationError) GetErrors() map[string]string {

	if a.Errors != nil {
		return a.Errors
	}
	return nil
}

func (a ApplicationError) GetMessage() string {
	if len(a.Message) > 0 {
		return a.Message
	}
	return a.Title
}

func (a ApplicationError) Error() string {
	return fmt.Sprintf("%v (%v) \n%v", a.Title, a.Code, a.Message)
}

var (
	InternalServerError ApplicationError = ApplicationError{
		Title: "Internal Server Error",
		Code:  http.StatusInternalServerError,
	}

	MethodNotAllowed ApplicationError = ApplicationError{
		Title: "Method Not Allowed",
		Code:  http.StatusMethodNotAllowed,
	}

	UserNotAuthenticated ApplicationError = ApplicationError{
		Title: "User Not Authenticated",
		Code:  http.StatusUnauthorized,
	}
	NotAuthorized ApplicationError = ApplicationError{
		Title: "User Not Authorized",
		Code:  http.StatusUnauthorized,
	}

	UserForbidden ApplicationError = ApplicationError{
		Title: "Authorization Error",
		Code:  http.StatusForbidden,
	}

	RequestValidationError ApplicationError = ApplicationError{
		Title: "Unprocessable Entity",
		Code:  http.StatusUnprocessableEntity,
	}

	BadRequest ApplicationError = ApplicationError{
		Title: "Bad Request",
		Code:  http.StatusBadRequest,
	}

	NotFound ApplicationError = ApplicationError{
		Title: "Not Found",
		Code:  http.StatusNotFound,
	}

	Conflict ApplicationError = ApplicationError{
		Title: "409 Conflict",
		Code:  http.StatusConflict,
	}
)

func (a ApplicationError) WithMessage(msg any, args ...any) *ApplicationError {
	switch msg := msg.(type) {
	case error:
		a.Message = msg.Error()
	case string:
		a.Message = fmt.Sprintf(msg, args...)
	}
	return &a
}

func (a *ApplicationError) WithErrors(errormap map[string]string) ApplicationError {
	a.Errors = errormap
	return *a
}

// TODO: Message ve Error eklemek icin bir method gerekebilir
