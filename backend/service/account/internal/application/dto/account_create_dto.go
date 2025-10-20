package dto

import (
	"strings"

	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type AccountCreateDTO struct {
	Username  string `json:"user_name"`
	AvatarUrl string `json:"avatar_url,omitempty"`
}

func (d AccountCreateDTO) Validate() error {
	var errors map[string]string = map[string]string{}

	if strings.EqualFold("", d.Username) || strings.Trim(d.Username, " ") == "" {
		errors["user_name"] = "Username cannot be empty"
	}

	if len(errors) != 0 {
		return coredomain.RequestValidationError.WithErrors(errors)
	}

	return nil
}
