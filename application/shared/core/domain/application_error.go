package domain

import (
	"fmt"
	"net/http"
)

type IApplicationError interface {
	GetMessage() string
	GetCode() int
}

type ApplicationError struct {
	Message string
	Code    int
}

func (a *ApplicationError) GetMessage() string {
	return a.Message
}

func (a *ApplicationError) GetCode() int {
	return a.Code
}

func (a *ApplicationError) Error() string {
	return fmt.Sprintf("Application Error with\nMessage: %v\nStatus: %v", a.Message, a.Code)
}

var InternalServerError ApplicationError = ApplicationError{
	Message: "Internal Server Error",
	Code:    http.StatusInternalServerError,
}

var MethodNotAllowed ApplicationError = ApplicationError{
	Message: "Method Not Allowed",
	Code:    http.StatusMethodNotAllowed,
}

var UserNotAuthenticated ApplicationError = ApplicationError{
	Message: "User Not Authenticated",
	Code:    http.StatusUnauthorized,
}

var UserForbidden ApplicationError = ApplicationError{
	Message: "Authorization Error",
	Code:    http.StatusForbidden,
}
