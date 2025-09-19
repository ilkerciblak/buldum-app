package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
)

type AccountRepository interface {
	GetById(ctx context.Context, userId uuid.UUID) (*model.Profile, error)
	GetAll(ctx context.Context) ([]*model.Profile, error)
	Create(ctx context.Context, p *model.Profile) error
	Update(ctx context.Context, userId uuid.UUID, p *model.Profile) error
	Delete(ctx context.Context, userId uuid.UUID) error
	Archive(ctx context.Context, userId uuid.UUID) error
}
