package mock

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/application"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

// type AccountRepository interface {
// 	GetById(ctx context.Context, userId uuid.UUID) (*model.Profile, error)
// 	GetAll(ctx context.Context) ([]*model.Profile, error)
// 	Create(ctx context.Context, p *model.Profile) error
// 	Update(ctx context.Context, userId uuid.UUID, p *model.Profile) error
// 	Delete(ctx context.Context, userId uuid.UUID) error
// 	Archive(ctx context.Context, userId uuid.UUID) error
// }

type MockAccountRepository struct {
}

func (m *MockAccountRepository) GetById(ctx context.Context, userId uuid.UUID) (*model.Profile, error) {
	if userId == uuid.Nil || userId == uuid.Max {
		return nil, coredomain.NotFound
	}

	return model.NewProfile("ilkerciblak", "url"), nil
}
func (m *MockAccountRepository) GetAll(ctx context.Context, params application.CommonQueryParameters, filter repository.ProfileGetAllQueryFilter) ([]*model.Profile, error) {

	return []*model.Profile{
		model.NewProfile("ilkerciblak", "url"),
		model.NewProfile("ilkerciblak", "url"),
	}, nil
}
func (m *MockAccountRepository) Create(ctx context.Context, p *model.Profile) error {
	if strings.EqualFold("error", p.Username) {
		return coredomain.BadRequest
	}

	return nil
}
func (m *MockAccountRepository) Update(ctx context.Context, userId uuid.UUID, p *model.Profile) error {
	return nil
}
func (m *MockAccountRepository) Delete(ctx context.Context, userId uuid.UUID) error {
	return nil
}
func (m *MockAccountRepository) Archive(ctx context.Context, userId uuid.UUID) error {
	if userId == uuid.Max {
		return coredomain.BadRequest.WithMessage("User Already Archived")
	}

	return nil
}
