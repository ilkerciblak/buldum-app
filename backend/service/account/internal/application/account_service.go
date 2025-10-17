package application

import (
	"context"
	"database/sql"

	"github.com/ilkerciblak/buldum-app/service/account/internal/application/command"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/query"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	"github.com/ilkerciblak/buldum-app/shared/logging"
)

type AccountServiceInterface interface {
	CreateAccount(createAccountCommand command.CreateAccountCommand, ctx context.Context) error
	UpdateAccount(updateAccountCommand command.UpdateAccountCommand, ctx context.Context) error
	ArchiveAccount(archiveAccountCommand command.ArchiveAccountCommand, ctx context.Context) error
	GetAccountById(query query.AccountGetByIdQuery, ctx context.Context) (*model.Profile, error)
	GetAllAccount(query query.AccountGetAllQuery, ctx context.Context) ([]*model.Profile, error)
}

type accountService struct {
	AccountRepository repository.IAccountRepository
	Logger            logging.ILogger
}

func AccountService(r repository.IAccountRepository, logger logging.ILogger) *accountService {
	return &accountService{
		AccountRepository: r,
		Logger:            logger,
	}
}

// Implement the IAccountService interface methods below:

func (s *accountService) CreateAccount(c command.CreateAccountCommand, ctx context.Context) error {
	if _, err := c.Validate(); err != nil {
		return err
	}

	// TODO: Check for conflicts
	account := model.NewProfile(c.Username, c.AvatarUrl)

	if err := s.AccountRepository.Create(ctx, account); err != nil {
		return coredomain.InternalServerError.WithMessage(err)
	}

	return nil

}

func (s *accountService) UpdateAccount(c command.UpdateAccountCommand, ctx context.Context) error {
	if _, err := c.Validate(); err != nil {
		return err.(coredomain.IApplicationError)
	}

	user, err := s.AccountRepository.GetById(ctx, c.UserId)
	if err != nil {
		return coredomain.BadRequest.WithMessage(err)
	}

	updated, err := user.UpdateProfile(model.UpdateUsername(c.Username), model.UpdateAvatarUrl(c.AvatarUrl))
	if err != nil {
		return coredomain.BadRequest.WithMessage(err)
	}

	if err := s.AccountRepository.Update(ctx, updated.Id, updated); err != nil {
		return coredomain.InternalServerError.WithMessage(err)
	}

	return nil

}

func (s *accountService) ArchiveAccount(c command.ArchiveAccountCommand, ctx context.Context) error {
	data, err := s.AccountRepository.GetById(ctx, c.Id)
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

	repoError := s.AccountRepository.Archive(ctx, archived.Id)
	if repoError != nil {
		return coredomain.InternalServerError.WithMessage(err)
	}

	return nil
}

func (s *accountService) GetAccountById(query query.AccountGetByIdQuery, ctx context.Context) (*model.Profile, error) {
	data, err := s.AccountRepository.GetById(ctx, query.Id)
	if err != nil {
		return nil, err.(coredomain.IApplicationError)
	}

	return data, nil
}

func (s *accountService) GetAllAccount(query query.AccountGetAllQuery, ctx context.Context) ([]*model.Profile, error) {
	data, err := s.AccountRepository.GetAll(ctx, query.CommonQueryParameters, query.ProfileGetAllQueryFilter)

	if err != nil {
		return nil, coredomain.BadRequest.WithMessage(err.Error())
	}

	return data, nil
}
