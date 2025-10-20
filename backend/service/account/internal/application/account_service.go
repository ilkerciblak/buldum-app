package application

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/dto"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	"github.com/ilkerciblak/buldum-app/shared/logging"
)

type accountService struct {
	AccountRepository repository.AccountRepositoryInterface
	Logger            logging.ILogger
}

func AccountService(r repository.AccountRepositoryInterface, logger logging.ILogger) *accountService {

	logger.With("Service", "Account-Service")

	return &accountService{
		AccountRepository: r,
		Logger:            logger,
	}
}

func (s *accountService) CreateAccount(dto dto.AccountCreateDTO, ctx context.Context) error {
	if count, err := s.AccountRepository.CountMatchingProfiles(ctx, dto.Username); err != nil {
		return err
	} else if count != 0 {
		return coredomain.Conflict.WithMessage("Username is already taken")
	}

	account := model.NewProfile(dto.Username, dto.AvatarUrl)

	if err := s.AccountRepository.Create(ctx, account); err != nil {
		return coredomain.InternalServerError.WithMessage(err)
	}

	return nil

}

func (s *accountService) UpdateAccount(userId uuid.UUID, dto dto.AccountUpdateDTO, ctx context.Context) error {

	user, err := s.AccountRepository.GetById(ctx, userId)
	if err != nil {
		return coredomain.BadRequest.WithMessage(err)
	}

	updated, err := user.UpdateProfile(model.UpdateUsername(dto.Username), model.UpdateAvatarUrl(dto.AvatarUrl))
	if err != nil {
		return coredomain.BadRequest.WithMessage(err)
	}

	if err := s.AccountRepository.Update(ctx, userId, updated); err != nil {
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

func (s *accountService) GetAccountById(userId uuid.UUID, ctx context.Context) (*dto.AccountResultDTO, error) {
	data, err := s.AccountRepository.GetById(ctx, userId)
	if err != nil {
		return nil, err.(coredomain.IApplicationError)
	}

	return dto.FromAccountModel(data), nil
}

func (s *accountService) GetAllAccount(query coredomain.CommonQueryParameters, filter repository.ProfileGetAllQueryFilter, ctx context.Context) ([]*dto.AccountResultDTO, error) {
	data, err := s.AccountRepository.GetAll(ctx, query, filter)

	if err != nil {
		return nil, coredomain.BadRequest.WithMessage(err.Error())
	}

	return dto.FromModelList(data), nil
}
