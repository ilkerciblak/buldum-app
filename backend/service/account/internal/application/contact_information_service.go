package application

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/dto"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/authentication"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	"github.com/ilkerciblak/buldum-app/shared/logging"
)

type ContactInformationService struct {
	repo     repository.ContactInformationRepositoryInterface
	userRepo repository.AccountRepositoryInterface
	logger   logging.ILogger
}

func NewContactInformationService(
	repo repository.ContactInformationRepositoryInterface,
	userRepo repository.AccountRepositoryInterface,
	logger logging.ILogger,
) *ContactInformationService {
	return &ContactInformationService{
		repo:     repo,
		userRepo: userRepo,
		logger:   logger,
	}
}

func (c *ContactInformationService) Create(dto dto.ContactInformationCreateDTO, ctx context.Context) error {
	// Validate if user exists
	_, err := c.userRepo.GetById(ctx, dto.AccountID)
	if err != nil {
		return err
	}

	//validate requesting user mathes with instance's user or requesting user is a admin roled whom
	userClaims, k := ctx.Value(authentication.AuthKey{}).(*authentication.UserClaims)
	if !k {
		return coredomain.UserNotAuthenticated
	}

	if userClaims.ID != dto.AccountID && !strings.EqualFold(userClaims.Role, "admin") {
		return coredomain.NotAuthorized
	}

	// create instance from dto
	m := model.NewContactInformation(
		dto.AccountID,
		model.ContactInformationTypeFromString(dto.Type),
		dto.Publicity,
		dto.ContactInfo,
	)

	// implement repository action
	if err := c.repo.CreateCI(*m, ctx); err != nil {
		return coredomain.InternalServerError.WithMessage(err)
	}

	return nil
}

func (c *ContactInformationService) Update(id uuid.UUID, dto dto.ContactInformationUpdateDTO, ctx context.Context) error {
	// check if user exists
	_, err := c.userRepo.GetById(ctx, dto.AccountID)
	if err != nil {
		return err
	}

	// check if instance exists
	m, err := c.repo.GetById(id, ctx)
	if err != nil {
		return err
	}

	// validate requesting user matches with instance's user or requesting user is a admin roled
	claims, k := ctx.Value(authentication.AuthKey{}).(*authentication.UserClaims)
	if !k {
		return coredomain.NotAuthorized
	}
	if claims.ID != m.UserID && !strings.EqualFold(claims.Role, "admin") {
		return coredomain.NotAuthorized
	}

	if err := m.BeforeUpdate(); err != nil {
		return err
	}
	if err := m.SetContactInfo(dto.ContactInfo); err != nil {
		return err
	}
	m.SetPublicity(dto.Publicity)

	if err := c.repo.UpdateCI(id, *m, ctx); err != nil {
		return err
	}

	return nil
}

func (c *ContactInformationService) Archive(id uuid.UUID, ctx context.Context) error {

	// check if instance exists
	m, err := c.repo.GetById(id, ctx)
	if err != nil {
		return err
	}

	// validate requesting user matches with instance's user or requesting user is a admin roled
	claims, k := ctx.Value(authentication.AuthKey{}).(*authentication.UserClaims)
	if !k {
		return coredomain.NotAuthorized
	}
	if claims.ID != m.UserID && !strings.EqualFold(claims.Role, "admin") {
		return coredomain.NotAuthorized
	}

	if err := m.BeforeUpdate(); err != nil {
		return err
	}
	m.Archive()

	if err := c.repo.UpdateCI(id, *m, ctx); err != nil {
		return err
	}

	return nil
}

func (c *ContactInformationService) GetAllByAccount(accountId uuid.UUID, ctx context.Context) ([]*dto.ContactInformationGetAllResultDTO, error) {
	var filter repository.ContactInformationQueryFilter
	filter.SetUserID(accountId)
	data, err := c.repo.GetAll(filter, ctx)
	if err != nil {
		return nil, err
	}

	return dto.FromModelListToList(data), nil
}

func (c *ContactInformationService) GetAllFiltered(filter dto.ContactInformationGetAllDTO, ctx context.Context) ([]*dto.ContactInformationGetAllResultDTO, error) {

	data, err := c.repo.GetAll(repository.ContactInformationQueryFilter{
		UserID:     filter.UserID,
		Type:       filter.Type,
		Publicity:  filter.Publicity,
		IsArchived: filter.IsArchived,
	},
		ctx,
	)
	if err != nil {
		return nil, err
	}
	return dto.FromModelListToList(data), nil
}
