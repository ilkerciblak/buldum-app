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
	return fmt.Sprintf("%v (%v)", a.Title, a.Code)
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
)

func (a ApplicationError) WithMessage(msg string) *ApplicationError {
	a.Message = msg
	return &a
}

func (a *ApplicationError) WithErrors(errormap map[string]string) ApplicationError {
	a.Errors = errormap
	return *a
}

// TODO: Message ve Error eklemek icin bir method gerekebilir
