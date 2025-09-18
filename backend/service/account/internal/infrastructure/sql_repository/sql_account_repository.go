package sqlrepository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	account_db "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql"
)

type SqlAccountRepository struct {
	Db account_db.Queries
}

func NewSqlAccountRepository(db account_db.Queries) *SqlAccountRepository {
	return &SqlAccountRepository{
		Db: db,
	}
}

func (s SqlAccountRepository) GetById(ctx context.Context, userId uuid.UUID) (*model.Profile, error) {
	return nil, nil
}

func (s SqlAccountRepository) GetAll(ctx context.Context) ([]*model.Profile, error) {
	return []*model.Profile{}, nil
}
func (s SqlAccountRepository) Create(ctx context.Context, p *model.Profile) error {
	if err := s.Db.CreateProfile(ctx, account_db.CreateProfileParams{
		ID:         p.Id,
		UserID:     p.UserID,
		UserName:   p.Username,
		AvatarUrl:  sql.NullString{String: p.AvatarUrl},
		CreatedAt:  p.CreatedAt,
		IsArchived: sql.NullBool{Bool: p.IsArchived},
	}); err != nil {
		return err
	}

	return nil
}

func (s SqlAccountRepository) Update(ctx context.Context, userId uuid.UUID, p *model.Profile) error {
	return nil
}
func (s SqlAccountRepository) Delete(ctx context.Context, userId uuid.UUID) error {
	return nil
}
