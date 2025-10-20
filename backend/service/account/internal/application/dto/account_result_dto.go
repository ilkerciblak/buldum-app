package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
)

type AccountResultDTO struct {
	Id         uuid.UUID `json:"id"`
	Username   string    `json:"user_name"`
	AvatarUrl  string    `json:"avatar_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	IsArchived bool      `json:"is_archived"`
}

func FromAccountModel(m *model.Profile) *AccountResultDTO {
	return &AccountResultDTO{
		Id:         m.Id,
		Username:   m.Username,
		AvatarUrl:  m.AvatarUrl,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
		DeletedAt:  m.DeletedAt,
		IsArchived: m.IsArchived,
	}
}

func FromModelList(ml []*model.Profile) []*AccountResultDTO {
	res := make([]*AccountResultDTO, len(ml))
	for i, m := range ml {
		res[i] = FromAccountModel(m)
	}
	return res
}
