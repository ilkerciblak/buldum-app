package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	Id         uuid.UUID
	Username   string
	AvatarUrl  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	IsArchived bool
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

type withUpdateFunc func(p *Profile) error

func UpdateUsername(val string) withUpdateFunc {
	return func(p *Profile) error {
		p.Username = val
		return nil
	}
}

func UpdateAvatarUrl(val string) withUpdateFunc {
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

func (p *Profile) UpdateProfile(fs ...withUpdateFunc) (*Profile, error) {
	for _, f := range fs {
		err := f(p)
		if err != nil {
			return p, err
		}
	}
	p.UpdatedAt = time.Now()

	return p, nil
}
