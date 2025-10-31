package mock

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

//	type AccountRepository interface {
//		GetById(ctx context.Context, userId uuid.UUID) (*model.Profile, error)
//		GetAll(ctx context.Context) ([]*model.Profile, error)
//		Create(ctx context.Context, p *model.Profile) error
//		Update(ctx context.Context, userId uuid.UUID, p *model.Profile) error
//		Delete(ctx context.Context, userId uuid.UUID) error
//		Archive(ctx context.Context, userId uuid.UUID) error
//	}
var now time.Time = time.Now()

var Id1 uuid.UUID = uuid.MustParse("370907d3-698d-40af-a1ce-c23ce40735c5")
var Id2 uuid.UUID = uuid.MustParse("2677a213-a037-409c-8f7b-21810eefe5de")
var Id3 uuid.UUID = uuid.MustParse("a5527cd3-2418-4415-912d-365e86048338")
var Id4 uuid.UUID = uuid.MustParse("44444444-4444-4444-4444-444444444444")

var AccountList map[uuid.UUID]*model.Profile = map[uuid.UUID]*model.Profile{
	Id1: {
		Id:         Id1,
		Username:   "ilkerciblak",
		AvatarUrl:  "url",
		CreatedAt:  now,
		IsArchived: false,
	},
	Id2: {
		Id:         Id2,
		Username:   "turkantugcetufan",
		AvatarUrl:  "url",
		CreatedAt:  now.Add(2 * time.Second),
		IsArchived: true,
	},
	Id3: {
		Id:         Id3,
		Username:   "luffy-chan",
		AvatarUrl:  "url",
		CreatedAt:  now.Add(3 * time.Second),
		IsArchived: false,
	},
}

type MockAccountRepository struct {
}

func (m MockAccountRepository) GetById(ctx context.Context, userId uuid.UUID) (*model.Profile, error) {

	model, exists := AccountList[userId]
	if !exists {
		return nil, coredomain.NotFound
	}

	return model, nil

}
func (m MockAccountRepository) GetAll(ctx context.Context, params coredomain.CommonQueryParameters, filter repository.ProfileGetAllQueryFilter) ([]*model.Profile, error) {
	res := []*model.Profile{}
	for _, account := range AccountList {

		if (filter.IsArchived == false || filter.IsArchived == account.IsArchived) &&
			(filter.Username == "" || strings.Contains(account.Username, filter.Username)) &&
			params.Limit > len(res) {
			res = append(res, account)
		}
	}

	return res, nil
}
func (m MockAccountRepository) Create(ctx context.Context, p *model.Profile) error {
	if p.Username == "" || p.Id == uuid.Max {

		return coredomain.BadRequest
	}

	if _, exists := AccountList[p.Id]; exists {
		return coredomain.ApplicationError{
			Code:    http.StatusConflict,
			Message: "Account With Id Already Registered",
		}
	}

	AccountList[p.Id] = p

	return nil
}

func (m MockAccountRepository) CountMatchingProfiles(ctx context.Context, username string) (int64, error) {
	for _, acc := range AccountList {
		if strings.EqualFold(acc.Username, username) {
			return 1, nil
		}
	}

	return 0, nil
}

func (m MockAccountRepository) Update(ctx context.Context, userId uuid.UUID, p *model.Profile) error {

	if _, exists := AccountList[userId]; exists {
		AccountList[userId] = p
	}

	return coredomain.NotFound
}
func (m MockAccountRepository) Delete(ctx context.Context, userId uuid.UUID) error {

	if acc, exists := AccountList[userId]; exists {
		updated, err := acc.UpdateProfile(model.ArchiveProfile)
		if err != nil {
			return err
		}
		AccountList[userId] = updated
	}

	return coredomain.NotFound
}
func (m MockAccountRepository) Archive(ctx context.Context, userId uuid.UUID) error {
	if acc, exists := AccountList[userId]; exists {
		updated, err := acc.UpdateProfile(model.ArchiveProfile)
		if err != nil {
			return err
		}
		AccountList[userId] = updated
	}

	return coredomain.NotFound
}
