package command

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type ArchiveAccountCommand struct {
	Id uuid.UUID `path:"id"`
}

func NewArchiveAccountCommand(id string) (*ArchiveAccountCommand, error) {
	userid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return &ArchiveAccountCommand{
		Id: userid,
	}, nil
}

func (c ArchiveAccountCommand) Handler(r repository.IAccountRepository, ctx context.Context) coredomain.IApplicationError {

	data, err := r.GetById(ctx, c.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return coredomain.NotFound.WithMessage(err)
		}

		return coredomain.InternalServerError.WithMessage(err)
	}

	archived, err := data.UpdateProfile(model.ArchiveProfile)
	if err != nil {
		return coredomain.BadRequest.WithMessage(err)
	}

	repoError := r.Archive(ctx, archived.Id)
	if repoError != nil {
		return coredomain.InternalServerError.WithMessage(err)
	}

	return nil

}
