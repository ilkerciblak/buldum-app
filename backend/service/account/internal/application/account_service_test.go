package application_test

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application"
	a_application "github.com/ilkerciblak/buldum-app/service/account/internal/application"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/dto"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	"github.com/ilkerciblak/buldum-app/shared/logging"
)

var now time.Time = time.Now()

var id1 uuid.UUID = uuid.MustParse("370907d3-698d-40af-a1ce-c23ce40735c5")
var id2 uuid.UUID = uuid.MustParse("2677a213-a037-409c-8f7b-21810eefe5de")
var id3 uuid.UUID = uuid.MustParse("a5527cd3-2418-4415-912d-365e86048338")

var accountList map[uuid.UUID]*model.Profile = map[uuid.UUID]*model.Profile{
	id1: {
		Id:         id1,
		Username:   "ilkerciblak",
		AvatarUrl:  "url",
		CreatedAt:  now,
		IsArchived: false,
	},
	id2: {
		Id:         id2,
		Username:   "turkantugcetufan",
		AvatarUrl:  "url",
		CreatedAt:  now.Add(2 * time.Second),
		IsArchived: true,
	},
	id3: {
		Id:         id3,
		Username:   "luffy-chan",
		AvatarUrl:  "url",
		CreatedAt:  now.Add(3 * time.Second),
		IsArchived: false,
	},
}

var accountService application.AccountServiceInterface = a_application.AccountService(
	MockAccountRepository{},
	logging.NewSlogger(
		logging.LoggerOptions{
			MinLevel:    logging.DEBUG,
			JsonLogging: true,
			LoggingRate: 1,
		},
	),
)

type MockAccountRepository struct {
}

func (m MockAccountRepository) GetById(ctx context.Context, userId uuid.UUID) (*model.Profile, error) {

	model, exists := accountList[userId]
	if !exists {
		return nil, coredomain.NotFound
	}

	return model, nil

}
func (m MockAccountRepository) GetAll(ctx context.Context, params coredomain.CommonQueryParameters, filter repository.ProfileGetAllQueryFilter) ([]*model.Profile, error) {
	res := []*model.Profile{}
	for _, account := range accountList {

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

	if _, exists := accountList[p.Id]; exists {
		return coredomain.ApplicationError{
			Code:    http.StatusConflict,
			Message: "Account With Id Already Registered",
		}
	}

	accountList[p.Id] = p

	return nil
}

func (m MockAccountRepository) CountMatchingProfiles(ctx context.Context, username string) (int64, error) {
	for _, acc := range accountList {
		if strings.EqualFold(acc.Username, username) {
			return 1, nil
		}
	}

	return 0, nil
}

func (m MockAccountRepository) Update(ctx context.Context, userId uuid.UUID, p *model.Profile) error {

	if _, exists := accountList[userId]; exists {
		accountList[userId] = p
	}

	return coredomain.NotFound
}
func (m MockAccountRepository) Delete(ctx context.Context, userId uuid.UUID) error {

	if acc, exists := accountList[userId]; exists {
		updated, err := acc.UpdateProfile(model.ArchiveProfile)
		if err != nil {
			return err
		}
		accountList[userId] = updated
	}

	return coredomain.NotFound
}
func (m MockAccountRepository) Archive(ctx context.Context, userId uuid.UUID) error {
	if acc, exists := accountList[userId]; exists {
		updated, err := acc.UpdateProfile(model.ArchiveProfile)
		if err != nil {
			return err
		}
		accountList[userId] = updated
	}

	return coredomain.NotFound
}

func TestApplicationLayer__TestGetAllAccount(t *testing.T) {

	// initiate getAllQuery with CommonQueryParameters and FilteringParameters

	cases := []struct {
		Name           string
		Query          coredomain.CommonQueryParameters
		Filter         repository.ProfileGetAllQueryFilter
		ExpectedResult struct {
			dataLength int
			err        error
		}
		DoesExpectsError bool
	}{
		{
			Name:   "Get All With No Query should Return 200 With All Values",
			Query:  *coredomain.NewCommonQueryParameters(),
			Filter: *repository.DefaultAccountGetAllQueryFilter(),
			ExpectedResult: struct {
				dataLength int
				err        error
			}{
				dataLength: len(accountList),
				err:        nil,
			},
			DoesExpectsError: false,
		},
		{
			Name:   "Get All With Using Limiting Should 200 With Limited Values",
			Query:  *coredomain.NewCommonQueryParameters(coredomain.SetLimit("1")),
			Filter: *repository.DefaultAccountGetAllQueryFilter(),
			ExpectedResult: struct {
				dataLength int
				err        error
			}{
				dataLength: 1,
				err:        nil,
			},
			DoesExpectsError: false,
		},
		{
			Name:  "Get All With Using Filtering Should 200 OK with Some Values",
			Query: *coredomain.NewCommonQueryParameters(),
			Filter: repository.ProfileGetAllQueryFilter{
				Username: "ilker",
			},
			ExpectedResult: struct {
				dataLength int
				err        error
			}{
				dataLength: 1,
				err:        nil,
			},
			DoesExpectsError: false,
		},
	}

	for _, tc := range cases {
		t.Run(
			tc.Name,
			func(t *testing.T) {
				data, err := accountService.GetAllAccount(tc.Query, tc.Filter, context.Background())

				if tc.DoesExpectsError {
					if err == nil {
						t.Fatalf("Error Expectations are Not Full-Filled")
					}
					if err.(coredomain.ApplicationError).GetCode() != tc.ExpectedResult.err.(coredomain.IApplicationError).GetCode() {
						t.Fatalf("Error Expectations are Not Full-Filled\nExpected %v\nGot %v", tc.ExpectedResult.err, err)
					}
				}

				if tc.ExpectedResult.dataLength != len(data) {
					t.Log(data)
					t.Fatalf("Expected Result with %v data\nGot %v", tc.ExpectedResult.dataLength, len(data))
				}

			},
		)
	}
}

func TestApplicationLayer__TestGetById(t *testing.T) {
	cases := []struct {
		Name           string
		Input          uuid.UUID
		ExpectedResult struct {
			ResultId uuid.UUID
			err      error
		}
		DoesExpectsError bool
	}{
		{
			Name:  "Should 200 OK With Related Value",
			Input: id1,
			ExpectedResult: struct {
				ResultId uuid.UUID
				err      error
			}{
				id1,
				nil,
			},
			DoesExpectsError: false,
		},
		{
			Name:  "Should 404 Not Found",
			Input: uuid.New(),
			ExpectedResult: struct {
				ResultId uuid.UUID
				err      error
			}{
				uuid.UUID{},
				coredomain.NotFound,
			},
			DoesExpectsError: true,
		},
	}

	for _, tc := range cases {
		t.Run(
			tc.Name,
			func(t *testing.T) {
				data, err := accountService.GetAccountById(tc.Input, context.Background())

				if tc.DoesExpectsError {
					if err == nil {
						t.Fatalf("Error Expectations are Not Full-Filled")
					}
					if err.(coredomain.ApplicationError).GetCode() != tc.ExpectedResult.err.(coredomain.IApplicationError).GetCode() {
						t.Fatalf("Error Expectations are Not Full-Filled\nExpected %v\nGot %v", tc.ExpectedResult.err, err)
					}
				}

				if !tc.DoesExpectsError && (tc.ExpectedResult.ResultId != data.Id) {
					t.Fatalf("Expected Id %v, Got %v", tc.ExpectedResult.ResultId, data.Id)
				}
			},
		)
	}
}

func TestApplicationLayer__TestCreateAccount(t *testing.T) {
	cases := []struct {
		Name             string
		Input            dto.AccountCreateDTO
		ExpectedResult   error
		DoesExpectsError bool
	}{
		{
			Name: "Should 201 Created",
			Input: dto.AccountCreateDTO{
				Username:  "avatareng",
				AvatarUrl: "url",
			},
			ExpectedResult:   nil,
			DoesExpectsError: false,
		},
		{
			Name: "Should 409 Conflict",
			Input: dto.AccountCreateDTO{
				Username: "ilkerciblak",
			},
			ExpectedResult:   coredomain.ApplicationError{Code: http.StatusConflict},
			DoesExpectsError: true,
		},
	}

	for _, tc := range cases {
		t.Run(
			tc.Name,
			func(t *testing.T) {
				countOfAccounts := len(accountList)
				err := accountService.CreateAccount(tc.Input, context.Background())

				if tc.DoesExpectsError {
					if err == nil {
						t.Fatalf("Error Expectations are Not Full-Filled")
					}
					if err.(*coredomain.ApplicationError).GetCode() != tc.ExpectedResult.(coredomain.IApplicationError).GetCode() {
						t.Fatalf("Error Expectations are Not Full-Filled\nExpected %v\nGot %v", tc.ExpectedResult, err)
					}
				}

				if !tc.DoesExpectsError && (len(accountList) == countOfAccounts) {
					t.Fatalf("Account Count not fullfills my precious expectations")
				}
			},
		)
	}
}
func TestApplicationLayer__TestArchiveAccount(t *testing.T) {
	cases := []struct {
		Name             string
		Input            uuid.UUID
		ExpectedResult   error
		DoesExpectsError bool
	}{
		{
			Name:             "Should OK",
			Input:            id1,
			ExpectedResult:   nil,
			DoesExpectsError: false,
		},
		{
			Name:             "Should Raise Error",
			Input:            id2,
			ExpectedResult:   coredomain.Conflict,
			DoesExpectsError: true,
		},
	}

	for _, tc := range cases {
		t.Run(
			tc.Name,
			func(t *testing.T) {
				err := accountService.ArchiveAccount(tc.Input, context.Background())

				if tc.DoesExpectsError {
					if err == nil {
						t.Fatalf("Error Expectations are Not Full-Filled")
					}
					if err.(*coredomain.ApplicationError).GetCode() != tc.ExpectedResult.(coredomain.IApplicationError).GetCode() {
						t.Fatalf("Error Expectations are Not Full-Filled\nExpected %v\nGot %v", tc.ExpectedResult, err)
					}
				}

				if !tc.DoesExpectsError && !accountList[tc.Input].IsArchived {
					t.Fatalf("Archive Command Not Archived the Given Object\n%v", accountList[tc.Input])
				}
			},
		)
	}
}
