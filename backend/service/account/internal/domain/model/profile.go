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

type WithUpdateFunc func(p *Profile) (*Profile, error)

func UpdateUsername(val string) WithUpdateFunc {
	return func(p *Profile) (*Profile, error) {
		p.Username = val
		return p, nil
	}
}

func UpdateAvatarUrl(val string) WithUpdateFunc {
	return func(p *Profile) (*Profile, error) {
		p.AvatarUrl = val
		return p, nil
	}
}

func ArchiveProfile(p *Profile) (*Profile, error) {
	if p.IsArchived {
		return p, fmt.Errorf("Profile Already Archived")
	}
	p.IsArchived = true
	return p, nil
}

func (p *Profile) UpdateProfile(fs ...WithUpdateFunc) (*Profile, error) {
	for _, f := range fs {
		pnew, err := f(p)
		if err != nil {
			return pnew, err
		}
		p = pnew
	}
	p.UpdatedAt = time.Now()

	return p, nil
}
