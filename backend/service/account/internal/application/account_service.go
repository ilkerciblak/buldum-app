package application

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	"github.com/ilkerciblak/buldum-app/shared/logging"
)

type accountService struct {
	AccountRepository repository.IAccountRepository
	Logger            logging.ILogger
}

func AccountService(r repository.IAccountRepository, logger logging.ILogger) *accountService {
	logger.With("Service", "Account-Service")
	return &accountService{
		AccountRepository: r,
		Logger:            logger,
	}
}

// Implement the IAccountService interface methods below:

func (s *accountService) CreateAccount(c model.Profile, ctx context.Context) error {

	// TODO: Check for conflicts
	account := model.NewProfile(c.Username, c.AvatarUrl)

	if err := s.AccountRepository.Create(ctx, account); err != nil {
		return coredomain.InternalServerError.WithMessage(err)
	}

	return nil

}

func (s *accountService) UpdateAccount(p model.Profile, ctx context.Context) error {

	user, err := s.AccountRepository.GetById(ctx, p.Id)
	if err != nil {
		return coredomain.BadRequest.WithMessage(err)
	}

	updated, err := user.UpdateProfile(model.UpdateUsername(p.Username), model.UpdateAvatarUrl(p.AvatarUrl))
	if err != nil {
		return coredomain.BadRequest.WithMessage(err)
	}

	if err := s.AccountRepository.Update(ctx, updated.Id, updated); err != nil {
		return coredomain.InternalServerError.WithMessage(err)
	}

	return nil

}

func (s *accountService) ArchiveAccount(userId uuid.UUID, ctx context.Context) error {
	data, err := s.AccountRepository.GetById(ctx, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return coredomain.NotFound.WithMessage(err)
		}

		return coredomain.InternalServerError.WithMessage(err)
	}
	// Already Archived
	if data.IsArchived {
		return coredomain.Conflict.WithMessage("Account is already archived")
	}

	archived, err := data.UpdateProfile(model.ArchiveProfile)
	if err != nil {
		return coredomain.BadRequest.WithMessage(err)
	}

	repoError := s.AccountRepository.Archive(ctx, archived.Id)
	if repoError != nil {
		return coredomain.InternalServerError.WithMessage(err)
	}

	return nil
}

func (s *accountService) GetAccountById(userId uuid.UUID, ctx context.Context) (*model.Profile, error) {
	data, err := s.AccountRepository.GetById(ctx, userId)
	if err != nil {
		return nil, err.(coredomain.IApplicationError)
	}

	return data, nil
}

func (s *accountService) GetAllAccount(query coredomain.CommonQueryParameters, filter repository.ProfileGetAllQueryFilter, ctx context.Context) ([]*model.Profile, error) {
	data, err := s.AccountRepository.GetAll(ctx, query, filter)

	if err != nil {
		return nil, coredomain.BadRequest.WithMessage(err.Error())
	}

	return data, nil
}
