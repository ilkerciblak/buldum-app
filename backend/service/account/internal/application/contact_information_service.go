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
	//TODO: authKey public ve shared kisma tasinacak

	//validate requesting user mathes with instance's user or requesting user is a admin roled whom
	type authKey struct{}
	userClaims := ctx.Value(authKey{}).(*authentication.UserClaims)

	if userClaims.ID != dto.AccountID && !strings.EqualFold(userClaims.Role, "admin") {
		return coredomain.NotAuthorized
	}

	// create instance from dto
	m := model.NewContactInformation(
		dto.AccountID,
		dto.Type,
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
	// Implementation here
	return nil
}

func (c *ContactInformationService) Archive(id uuid.UUID, ctx context.Context) error {

}

func (c *ContactInformationService) GetSingleByAccountByType(accountId uuid.UUID, ciType string, ctx context.Context) (*dto.ContactInformationGetAllByAccountResultDTO, error) {
	data, err := c.repo.GetByAccountAndType(accountId, ciType, ctx)
	if err != nil {
		return nil, err
	}

	// Implementation here
	return dto.FromModel(data), nil
}

func (c *ContactInformationService) GetAllByAccount(accountId uuid.UUID, ctx context.Context) []*dto.ContactInformationGetAllByAccountResultDTO {
	data := c.repo.GetAllByAccountId(accountId, ctx)

	return dto.FromModelListToList(data)
}
