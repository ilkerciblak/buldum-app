package command

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type UpdateAccountCommand struct {
	UserId    uuid.UUID `path:"user_id"`
	Username  string    `json:"user_name"`
	AvatarUrl string    `json:"avatar_url"`
}

func NewUpdateAccountCommand(m map[string]interface{}) (*UpdateAccountCommand, error) {
	userId, err := uuid.Parse(m["user_id"].(string))
	if err != nil {
		return nil, err
	}

	return &UpdateAccountCommand{
		UserId: userId,
	}, nil
}

func (c *UpdateAccountCommand) Validate() (*UpdateAccountCommand, error) {
	errors := map[string]string{}

	if len(strings.Trim(c.Username, " ")) <= 0 {
		errors["username"] = "Username field is required"
	}

	if len(errors) > 0 {
		return nil, coredomain.RequestValidationError.WithErrors(errors)
	}

	return c, nil
}

func (c *UpdateAccountCommand) Handler(r repository.IAccountRepository, ctx context.Context) coredomain.IApplicationError {
	if _, err := c.Validate(); err != nil {
		return err.(coredomain.IApplicationError)
	}

	user, err := r.GetById(ctx, c.UserId)
	if err != nil {
		return coredomain.BadRequest.WithMessage(err)
	}

	updated, err := user.UpdateProfile(model.UpdateUsername(c.Username), model.UpdateAvatarUrl(c.AvatarUrl))
	if err != nil {
		return coredomain.BadRequest.WithMessage(err)
	}

	if err := r.Update(ctx, updated.Id, updated); err != nil {
		return coredomain.InternalServerError.WithMessage(err)
	}

	return nil
}
