package sqlrepository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	account_db "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql"
	"github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql/mapper"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type SqlAccountRepository struct {
	Db account_db.Queries
	// Logger logging.ILogger
}

func NewSqlAccountRepository(db account_db.Queries) *SqlAccountRepository {
	return &SqlAccountRepository{
		Db: db,
	}
}

func (s SqlAccountRepository) GetById(ctx context.Context, userId uuid.UUID) (*model.Profile, error) {
	data, err := s.Db.GetProfileById(ctx, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, coredomain.NotFound.WithMessage("No User Found With userId %v", userId)
		}
		// TODO: BETTER ERROR HANDLING
		return nil, coredomain.InternalServerError.WithMessage("%v", err.Error())
	}

	return mapper.DBModelToDTO(data), nil
}

func (s SqlAccountRepository) GetAll(ctx context.Context, params coredomain.CommonQueryParameters, filter repository.ProfileGetAllQueryFilter) ([]*model.Profile, error) {

	data, err := s.Db.GetAllProfile(ctx, account_db.GetAllProfileParams{
		Column1:    params.Sort,
		Limit:      int32(params.Limit),
		Offset:     int32(params.Offset),
		Column4:    params.Order,
		UserName:   filter.Username,
		IsArchived: filter.IsArchived,
	})
	if err != nil {
		return nil, coredomain.InternalServerError.WithMessage(err)
	}

	return mapper.DBModelListToDTO(data), nil
}
func (s SqlAccountRepository) Create(ctx context.Context, p *model.Profile) error {
	if err := s.Db.CreateProfile(ctx, account_db.CreateProfileParams{
		ID:         p.Id,
		UserName:   p.Username,
		AvatarUrl:  sql.NullString{String: p.AvatarUrl, Valid: true},
		CreatedAt:  p.CreatedAt,
		IsArchived: sql.NullBool{Bool: p.IsArchived, Valid: true},
	}); err != nil {
		return err
	}

	return nil
}

func (s SqlAccountRepository) Update(ctx context.Context, userId uuid.UUID, p *model.Profile) error {
	upd := account_db.UpdateProfileParams{
		ID:        userId,
		UserName:  p.Username,
		AvatarUrl: sql.NullString{String: p.AvatarUrl, Valid: true},
		UpdatedAt: sql.NullTime{Time: p.UpdatedAt, Valid: true},
	}
	if err := s.Db.UpdateProfile(ctx, upd); err != nil {
		if err == sql.ErrNoRows {
			return coredomain.NotFound.WithMessage("No User Found With userId: %v", userId)
		}
		return coredomain.InternalServerError.WithMessage(err)
	}

	return nil

}
func (s SqlAccountRepository) Delete(ctx context.Context, userId uuid.UUID) error {
	if err := s.Db.DeleteProfile(ctx, userId); err != nil {

		if err == sql.ErrNoRows {
			return coredomain.NotFound.WithMessage("No User Found With userId: %v", userId)
		}
		return coredomain.InternalServerError.WithMessage(err)

	}
	return nil
}

func (s SqlAccountRepository) Archive(ctx context.Context, userId uuid.UUID) error {
	if err := s.Db.ArchiveProfile(ctx, userId); err != nil {

		if err == sql.ErrNoRows {
			return coredomain.NotFound.WithMessage("No User Found With userId: %v", userId)
		}
		return coredomain.InternalServerError.WithMessage(err)

	}
	return nil
}

func (s SqlAccountRepository) CountMatchingProfiles(ctx context.Context, username string) (int64, error) {
	res, err := s.Db.CountMatchingProfiles(ctx, username)
	if err != nil {
		return 0, coredomain.InternalServerError
	}

	return res, nil
}
