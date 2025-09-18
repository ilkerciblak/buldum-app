package command

import (
	"context"
	"strings"

	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type CreateAccountCommand struct {
	Username  string `json:"user_name"`
	AvatarUrl string `json:"avatar_url"`
}

func (c *CreateAccountCommand) Validate() (*CreateAccountCommand, coredomain.IApplicationError) {
	var errors map[string]string = map[string]string{}

	if len(strings.Trim(c.Username, " ")) <= 0 {
		errors["username"] = "Username field is required"
	}

	if len(errors) != 0 {
		return nil, coredomain.RequestValidationError.WithErrors(errors)
	}

	return c, nil
}

func (c *CreateAccountCommand) Handler(r repository.AccountRepository, ctx context.Context) coredomain.IApplicationError {
	// Validate the request
	if _, err := c.Validate(); err != nil {
		return err
	}

	account := model.NewProfile(c.Username, c.AvatarUrl)

	if err := r.Create(ctx, account); err != nil {
		return coredomain.MethodNotAllowed.WithMessage(err.Error())
	}

	return nil
}
