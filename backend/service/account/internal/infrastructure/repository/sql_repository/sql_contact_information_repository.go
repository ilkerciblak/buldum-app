package sqlrepository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	account_db "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

type SqlContactInformationRepository struct {
	db account_db.Queries
}

func NewSqlContactInformationRepository(db account_db.Queries) *SqlContactInformationRepository {
	return &SqlContactInformationRepository{
		db: db,
	}
}

func (r SqlContactInformationRepository) GetById(id uuid.UUID, ctx context.Context) (*model.ContactInformation, error) {
	data, err := r.db.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, coredomain.NotFound
		}
		return nil, err
	}

	return account_db.ContactInformationDBModelToModel(data), nil
}
func (r SqlContactInformationRepository) GetAll(filter repository.ContactInformationQueryFilter, ctx context.Context) ([]*model.ContactInformation, error) {
	data, err := r.db.GetAllContactInformation(
		ctx,
		account_db.GetAllContactInformationParams{
			Column1: filter.UserID,
			Column2: filter.Type,
			Column3: filter.Publicity,
			Column4: filter.IsArchived,
		},
	)
	if err != nil {
		return nil, err
	}

	return account_db.ContactInformationDbModelListToModelList(data), nil

}
func (r SqlContactInformationRepository) CreateCI(model model.ContactInformation, ctx context.Context) error {
	if err := r.db.CreateContactInformation(
		ctx,
		account_db.CreateContactInformationParams{
			ID:                     model.Id,
			ProfileID:              uuid.NullUUID{UUID: model.UserID, Valid: true},
			ContactInformationType: sql.NullString{String: model.Type.String(), Valid: true},
			IsPublic:               sql.NullBool{Bool: model.Publicity, Valid: true},
			ContactInformation:     sql.NullString{String: model.ContactInfo, Valid: true},
			CreatedAt:              sql.NullTime{Time: model.CreatedAt, Valid: true},
			IsArchived:             sql.NullBool{Bool: model.IsArchived, Valid: true},
		},
	); err != nil {
		return err
	}

	return nil
}
func (r SqlContactInformationRepository) UpdateCI(id uuid.UUID, updated model.ContactInformation, ctx context.Context) error {
	if err := r.db.UpdateContactInformation(
		ctx,
		account_db.UpdateContactInformationParams{
			ProfileID:              uuid.NullUUID{UUID: updated.UserID, Valid: true},
			ContactInformationType: sql.NullString{String: updated.Type.String(), Valid: true},
			IsPublic:               sql.NullBool{Bool: updated.Publicity, Valid: true},
			ContactInformation:     sql.NullString{String: updated.ContactInfo, Valid: true},
			UpdatedAt:              sql.NullTime{Time: updated.UpdatedAt, Valid: true},
		},
	); err != nil {
		return err
	}

	return nil
}
func (r SqlContactInformationRepository) ArchiveCI(id uuid.UUID, ctx context.Context) error {
	if err := r.db.ArchiveContactInformation(ctx, account_db.ArchiveContactInformationParams{
		ID:         id,
		IsArchived: sql.NullBool{Bool: true, Valid: true},
		DeletedAt:  sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt:  sql.NullTime{Time: time.Now(), Valid: true},
	}); err != nil {
		return err
	}
	return nil
}
