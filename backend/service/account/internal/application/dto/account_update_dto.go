package dto

import (
	"strings"

	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type AccountUpdateDTO struct {
	Username  string `json:"user_name"`
	AvatarUrl string `json:"avatar_url,omitempty"`
}

func (d AccountUpdateDTO) Validate() error {
	var errors map[string]string = map[string]string{}

	if strings.EqualFold(d.Username, "") || len(strings.Trim(d.Username, " ")) == 0 {
		errors["user_name"] = "Username field cannot be empty"
	}

	if len(errors) != 0 {
		return coredomain.RequestValidationError.WithErrors(errors)
	}

	return nil
}
