package application_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application"
	a_application "github.com/ilkerciblak/buldum-app/service/account/internal/application"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/dto"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/mock"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	"github.com/ilkerciblak/buldum-app/shared/logging"
)

var accountService application.AccountServiceInterface = a_application.AccountService(
	mock.MockAccountRepository{},
	logging.NewSlogger(
		logging.LoggerOptions{
			MinLevel:    logging.DEBUG,
			JsonLogging: true,
			LoggingRate: 1,
		},
	),
)

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
			Query:  *coredomain.DefaultCommonQueryParameters(),
			Filter: *repository.DefaultAccountGetAllQueryFilter(),
			ExpectedResult: struct {
				dataLength int
				err        error
			}{
				dataLength: len(mock.AccountList),
				err:        nil,
			},
			DoesExpectsError: false,
		},
		{
			Name: "Get All With Using Limiting Should 200 With Limited Values",
			Query: func() coredomain.CommonQueryParameters {
				q := *coredomain.DefaultCommonQueryParameters()
				q.SetPagination("1", "")
				return q

			}(),
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
			Query: *coredomain.DefaultCommonQueryParameters(),
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
			Input: mock.Id1,
			ExpectedResult: struct {
				ResultId uuid.UUID
				err      error
			}{
				mock.Id1,
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
				countOfAccounts := len(mock.AccountList)
				err := accountService.CreateAccount(tc.Input, context.Background())

				if tc.DoesExpectsError {
					if err == nil {
						t.Fatalf("Error Expectations are Not Full-Filled")
					}
					if err.(*coredomain.ApplicationError).GetCode() != tc.ExpectedResult.(coredomain.IApplicationError).GetCode() {
						t.Fatalf("Error Expectations are Not Full-Filled\nExpected %v\nGot %v", tc.ExpectedResult, err)
					}
				}

				if !tc.DoesExpectsError && (len(mock.AccountList) == countOfAccounts) {
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
			Input:            mock.Id1,
			ExpectedResult:   nil,
			DoesExpectsError: false,
		},
		{
			Name:             "Should Raise Error",
			Input:            mock.Id2,
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

				if !tc.DoesExpectsError && !mock.AccountList[tc.Input].IsArchived {
					t.Fatalf("Archive Command Not Archived the Given Object\n%v", mock.AccountList[tc.Input])
				}
			},
		)
	}
}
