package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	Id         uuid.UUID `json:"id"`
	Username   string    `json:"user_name"`
	AvatarUrl  string    `json:"avatar_url"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
	IsArchived bool      `json:"is_archived"`
}

func NewProfile(username, avatarUrl string) *Profile {
	return &Profile{
		Id:         uuid.Must(uuid.NewV7()),
		Username:   username,
		AvatarUrl:  avatarUrl,
		CreatedAt:  time.Now(),
		IsArchived: false,
	}
}

type WithUpdateFunc func(p *Profile) error

func UpdateUsername(val string) WithUpdateFunc {
	return func(p *Profile) error {
		p.Username = val
		return nil
	}
}

func UpdateAvatarUrl(val string) WithUpdateFunc {
	return func(p *Profile) error {
		p.AvatarUrl = val
		return nil
	}
}

func ArchiveProfile(p *Profile) error {
	if p.IsArchived {
		return fmt.Errorf("Profile Already Archived")
	}
	p.IsArchived = true
	return nil
}

func (p *Profile) UpdateProfile(fs ...WithUpdateFunc) (*Profile, error) {
	for _, f := range fs {
		err := f(p)
		if err != nil {
			return p, err
		}
	}
	p.UpdatedAt = time.Now()

	return p, nil
}
