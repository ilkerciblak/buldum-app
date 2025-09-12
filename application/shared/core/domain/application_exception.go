package domain

import (
	"fmt"
	"net/http"
)

type IApplicationException interface {
	GetTitle() string
	GetCode() int
	GetErrors() map[string]string
	GetMessage() string
}

type ApplicationException struct {
	// Title
	Code    int               `json:"code"`
	Title   string            `json:"title,omitempty"`
	Message string            `json:"message,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}

func (a *ApplicationException) GetTitle() string {
	return a.Title
}

func (a *ApplicationException) GetCode() int {
	return a.Code
}

func (a *ApplicationException) GetErrors() map[string]string {

	if a.Errors != nil {
		return a.Errors
	}
	return nil
}

func (a *ApplicationException) GetMessage() string {
	if len(a.Message) > 0 {
		return a.Message
	}
	return a.Title
}

func (a ApplicationException) Error() string {
	return fmt.Sprintf("%v (%v)", a.Title, a.Code)
}

var (
	InternalServerError ApplicationException = ApplicationException{
		Title: "Internal Server Error",
		Code:  http.StatusInternalServerError,
	}

	MethodNotAllowed ApplicationException = ApplicationException{
		Title: "Method Not Allowed",
		Code:  http.StatusMethodNotAllowed,
	}

	UserNotAuthenticated ApplicationException = ApplicationException{
		Title: "User Not Authenticated",
		Code:  http.StatusUnauthorized,
	}

	UserForbidden ApplicationException = ApplicationException{
		Title: "Authorization Error",
		Code:  http.StatusForbidden,
	}

	ValidationException ApplicationException = ApplicationException{
		Title: "Unprocessable Entity",
		Code:  http.StatusUnprocessableEntity,
	}
)

// func (a *ApplicationException) WithMessage(err error) *ApplicationException {
// 	return a
// }

// TODO: Message ve Error eklemek icin bir method gerekebilir
