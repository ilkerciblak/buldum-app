package mapper

import (
	"database/sql"

	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	account_db "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql"
)

func DBModelToDTO(a account_db.AccountProfile) *model.Profile {
	return &model.Profile{
		Id:         a.ID,
		Username:   a.UserName,
		AvatarUrl:  a.AvatarUrl.String,
		CreatedAt:  a.CreatedAt,
		UpdatedAt:  a.UpdatedAt.Time,
		DeletedAt:  a.DeletedAt.Time,
		IsArchived: a.IsArchived.Bool,
	}
}

func DBModelListToDTO(l []account_db.AccountProfile) []*model.Profile {
	ml := make([]*model.Profile, len(l))

	for _, dbm := range l {
		ml = append(ml, DBModelToDTO(dbm))
	}

	return ml
}

func DTOtoDBModel(m model.Profile) *account_db.AccountProfile {
	return &account_db.AccountProfile{
		ID:         m.Id,
		UserName:   m.Username,
		AvatarUrl:  sql.NullString{String: m.AvatarUrl},
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  sql.NullTime{Time: m.UpdatedAt},
		DeletedAt:  sql.NullTime{Time: m.DeletedAt},
		IsArchived: sql.NullBool{Bool: m.IsArchived},
	}
}
