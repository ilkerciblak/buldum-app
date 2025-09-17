package repository

import (
	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
)

type AccountRepository interface {
	GetById(userId uuid.UUID) (*model.Profile, error)
	GetAll() ([]*model.Profile, error)
	Create(p *model.Profile) error
	Update(userId uuid.UUID, p *model.Profile) error
	Delete(userId uuid.UUID) error
}
