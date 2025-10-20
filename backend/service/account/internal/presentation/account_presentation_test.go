package presentation_test

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/dto"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	"github.com/ilkerciblak/buldum-app/service/account/internal/presentation"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
	"github.com/ilkerciblak/buldum-app/shared/logging"
)

func TestEndPoint__UpdateAccount(t *testing.T) {
	var endpoint corepresentation.IEndPoint = presentation.UpdateAccountEndPoint{
		Service: &MockAccountService{},
	}

	cases := []struct {
		Name               string
		TestRequest        func() *http.Request
		ExpectedStatusCode int
	}{
		{
			Name: "Valid Update Request",
			TestRequest: func() *http.Request {
				return httptest.NewRequest(
					http.MethodPut,
					fmt.Sprintf("/accounts/%v", id1),
					bytes.NewBuffer([]byte(`{"user_name":"ilkerciblak"}`)),
				)
			},
			ExpectedStatusCode: http.StatusNoContent,
		},
		{
			Name: "Invalid Method",
			TestRequest: func() *http.Request {
				return httptest.NewRequest(
					http.MethodGet,
					fmt.Sprintf("/accounts/%v", uuid.New()),
					bytes.NewReader([]byte(`{"user_name":"ilkerciblak"}`)),
				)
			},
			ExpectedStatusCode: http.StatusMethodNotAllowed,
		},
		{
			Name: "Missing Body",
			TestRequest: func() *http.Request {
				return httptest.NewRequest(
					http.MethodPut,
					fmt.Sprintf("/accounts/%v", id2),
					nil,
				)
			},
			ExpectedStatusCode: http.StatusBadRequest,
		},
		{
			Name: "Malformed JSON Body",
			TestRequest: func() *http.Request {
				return httptest.NewRequest(http.MethodPut, fmt.Sprintf("/accounts/%v", id2), bytes.NewReader([]byte(`{invalid-json"`)))
			},

			ExpectedStatusCode: http.StatusBadRequest,
		},
		{
			Name: "422 Validation Error",
			TestRequest: func() *http.Request {
				return httptest.NewRequest(http.MethodPut, fmt.Sprintf("/accounts/%v", id2), bytes.NewReader([]byte(`{"avatar_url":"blabla"}`)))
			},
			ExpectedStatusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {
				req := c.TestRequest()
				mux := http.NewServeMux()
				testResponseWriter := httptest.NewRecorder()
				mux.HandleFunc(endpoint.Path(), corepresentation.GenerateHandlerFuncFromEndPoint(endpoint, &MockLogger{}))
				mux.ServeHTTP(testResponseWriter, req.Clone(t.Context()))

				if testResponseWriter.Result().StatusCode != c.ExpectedStatusCode {
					t.Fatalf("Status Code:\nExpected: %v\nGot %v", c.ExpectedStatusCode, testResponseWriter.Result().StatusCode)
				}

				if c.ExpectedStatusCode == http.StatusNoContent {
					if testResponseWriter.Body.Len() != 0 {
						t.Fatalf("Expected empty body for 204 No Content, got: %s",
							testResponseWriter.Body.String())
					}
				}

			},
		)

	}
}
func TestEndPoint__GetById(t *testing.T) {
	cases := []struct {
		Name            string
		Input           func() *http.Request
		DoesExpectError bool
		ExpectedOutput  corepresentation.ApiResult[*dto.AccountResultDTO]
		ExpectedError   coredomain.IApplicationError
	}{
		{
			Name:            "Invalid Request With POST, 405",
			Input:           func() *http.Request { return httptest.NewRequest("POST", fmt.Sprintf("/accounts/%v", id1), nil) },
			DoesExpectError: true,
			ExpectedOutput:  corepresentation.ApiResult[*dto.AccountResultDTO]{},
			ExpectedError:   coredomain.MethodNotAllowed,
		},
		{
			Name:            "Invalid Request With Invalid Id PathVal, 400 Bad Request",
			Input:           func() *http.Request { return httptest.NewRequest("GET", "/accounts/1", nil) },
			DoesExpectError: true,
			ExpectedOutput:  corepresentation.ApiResult[*dto.AccountResultDTO]{},
			ExpectedError:   coredomain.BadRequest,
		},
		{
			Name:            "Valid Request, 200 Ok",
			Input:           func() *http.Request { return httptest.NewRequest("GET", fmt.Sprintf("/accounts/%v", id2), nil) },
			DoesExpectError: false,
			ExpectedOutput: corepresentation.ApiResult[*dto.AccountResultDTO]{
				StatusCode: 200,
			},
			ExpectedError: nil,
		},
	}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {

				endPoint := presentation.AccountGetByIdEndPoint{
					Service: MockAccountService{},
				}

				req := c.Input()
				mux := http.NewServeMux()
				testResponseWriter := httptest.NewRecorder()
				mux.HandleFunc(endPoint.Path(), corepresentation.GenerateHandlerFuncFromEndPoint(endPoint, &MockLogger{}))
				mux.ServeHTTP(testResponseWriter, req.Clone(t.Context()))

				if c.DoesExpectError && testResponseWriter.Result().StatusCode != c.ExpectedError.GetCode() {
					t.Fatalf("Err Status Code:\nExpected: %v\nGot %v", c.ExpectedError.GetCode(), testResponseWriter.Result().StatusCode)
				}

				if !c.DoesExpectError && testResponseWriter.Result().StatusCode != c.ExpectedOutput.StatusCode {
					t.Fatalf("Status Code:\nExpected: %v\nGot %v", c.ExpectedOutput.StatusCode, testResponseWriter.Result().StatusCode)
				}

			},
		)
	}

}
func TestEndPoint__CreateAccountEndPoint(t *testing.T) {
	endPoint := presentation.CreateAccountEndPoint{
		Service: MockAccountService{},
	}

	cases := []struct {
		Name            string
		DoesExpectError bool
		TestRequest     *http.Request
		ExpectedError   coredomain.IApplicationError
	}{
		{
			Name:            "Create Account With user_name field only, VALID",
			DoesExpectError: false,
			TestRequest:     httptest.NewRequest("POST", "/accounts", bytes.NewReader([]byte(`{"user_name":"newnew"}`))),
			ExpectedError:   nil,
		},
		{
			Name:            "Create Account With wrong fields, INVALID with 422",
			DoesExpectError: true,
			TestRequest:     httptest.NewRequest("POST", "/accounts", bytes.NewReader([]byte(`{"username":"ilkerciblak"}`))),
			ExpectedError:   coredomain.RequestValidationError,
		},
		{
			Name:            "Create Account With wrong fields, INVALID with 400",
			DoesExpectError: true,
			TestRequest:     httptest.NewRequest("POST", "/accounts", bytes.NewReader([]byte(`{"user_name":"ilkerciblak}`))),
			ExpectedError:   coredomain.BadRequest,
		},
		{
			Name:            "Create Account With wrong fields, INVALID with 409 Conflict",
			DoesExpectError: true,
			TestRequest:     httptest.NewRequest("POST", "/accounts", bytes.NewReader([]byte(`{"user_name":"ilkerciblak}`))),
			ExpectedError:   coredomain.BadRequest,
		},
	}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {
				req := c.TestRequest
				mux := http.NewServeMux()
				testResponseWriter := httptest.NewRecorder()
				mux.HandleFunc(endPoint.Path(), corepresentation.GenerateHandlerFuncFromEndPoint(endPoint, &MockLogger{}))
				mux.ServeHTTP(testResponseWriter, req.Clone(t.Context()))

				if c.DoesExpectError && testResponseWriter.Result().StatusCode != c.ExpectedError.GetCode() {
					t.Fatalf("Err Status Code:\nExpected: %v\nGot %v", c.ExpectedError.GetCode(), testResponseWriter.Result().StatusCode)
				}

				if !c.DoesExpectError && testResponseWriter.Result().StatusCode != http.StatusCreated {
					t.Fatalf("Status Code:\nExpected: %v\nGot %v", http.StatusCreated, testResponseWriter.Result().StatusCode)
				}
			},
		)

	}
}
func TestEndPoint__ArchiveAccountEndPoint(t *testing.T) {
	archiveAccountEndPoint := presentation.ArchiveAccountEndPoint{
		Service: MockAccountService{},
	}

	mux := http.NewServeMux()
	mux.HandleFunc(archiveAccountEndPoint.Path(), corepresentation.GenerateHandlerFuncFromEndPoint(archiveAccountEndPoint, &MockLogger{}))

	cases := []struct {
		Name            string
		TestRequest     *http.Request
		DoesExpectError bool
		ExpectedError   coredomain.IApplicationError
	}{

		{
			Name:            "Given Request Should Return 400 Bad Request",
			TestRequest:     httptest.NewRequest("POST", fmt.Sprintf("/accounts/%d/archive", 12345), nil),
			DoesExpectError: true,
			ExpectedError:   coredomain.BadRequest,
		},
		{
			Name:            "Given Request Should Return 405 Method Not Allowed",
			TestRequest:     httptest.NewRequest("PUT", fmt.Sprintf("/accounts/%s/archive", uuid.New()), nil),
			DoesExpectError: true,
			ExpectedError:   coredomain.MethodNotAllowed,
		},
		{
			Name:            "Given Request Should Return 204 No Content",
			TestRequest:     httptest.NewRequest("POST", fmt.Sprintf("/accounts/%s/archive", id1), nil),
			DoesExpectError: false,
			ExpectedError:   nil,
		},
		{
			Name:            "Given Request Should Return 409 Conflict",
			TestRequest:     httptest.NewRequest("POST", fmt.Sprintf("/accounts/%s/archive", id2), nil),
			DoesExpectError: true,
			ExpectedError:   coredomain.Conflict,
		},
	}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {
				testResponseWriter := httptest.NewRecorder()
				mux.ServeHTTP(testResponseWriter, c.TestRequest)
				result := archiveAccountEndPoint.HandleRequest(testResponseWriter, c.TestRequest)

				if c.DoesExpectError {
					if result.Error == nil {
						t.Fatalf("Error Expectations was not satisfied")
					}

					if result.Error.(coredomain.IApplicationError).GetCode() != c.ExpectedError.GetCode() {
						t.Fatalf("Error Expectations was not satisfied\nGot %v\n Expected %v", result.Error, c.ExpectedError)
					}

				} else {
					if result.Error != nil {
						t.Fatalf("Test was not expecting error but\nGot %v\nExpected %v\n", result.Error, http.StatusCreated)
					}

					if result.Data != nil {
						t.Fatalf("Response Data is not as expected, Expected no data, Got %v", result.Data)
					}

					if result.StatusCode != http.StatusNoContent {
						t.Fatalf("Response Status Code is not as expected, Expects %d, Got %d", http.StatusCreated, result.StatusCode)
					}
				}
			},
		)
	}
}

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

type MockAccountService struct {
}

func (m MockAccountService) CreateAccount(dto dto.AccountCreateDTO, ctx context.Context) error {

	for _, acc := range accountList {
		if strings.EqualFold(acc.Username, dto.Username) {
			return coredomain.Conflict
		}
	}

	return nil
}
func (m MockAccountService) UpdateAccount(id uuid.UUID, dto dto.AccountUpdateDTO, ctx context.Context) error {

	_, exists := accountList[id]
	if !exists {
		return coredomain.NotFound
	}

	return nil

}
func (m MockAccountService) ArchiveAccount(userId uuid.UUID, ctx context.Context) error {
	mp, exists := accountList[userId]
	if !exists {
		return coredomain.NotFound
	}

	if mp.IsArchived {
		return coredomain.Conflict
	}

	return nil

}
func (m MockAccountService) GetAccountById(userId uuid.UUID, ctx context.Context) (*dto.AccountResultDTO, error) {
	model, exists := accountList[userId]
	if !exists {
		return nil, coredomain.NotFound
	}

	return dto.FromAccountModel(model), nil
}
func (m MockAccountService) GetAllAccount(query coredomain.CommonQueryParameters, filter repository.ProfileGetAllQueryFilter, ctx context.Context) ([]*dto.AccountResultDTO, error) {
	res := make([]*dto.AccountResultDTO, 0)
	for _, acc := range accountList {
		if query.Limit > len(res) {
			break
		}
		if (filter.IsArchived == false || filter.IsArchived == acc.IsArchived) &&
			(filter.Username == "" || strings.Contains(acc.Username, filter.Username)) {
			res = append(res, dto.FromAccountModel(acc))
			continue
		}

	}

	return res, nil
}

type MockLogger struct{}

func (m MockLogger) DEBUG(ctx context.Context, msg string, args ...interface{}) {}
func (m MockLogger) INFO(ctx context.Context, msg string, args ...interface{})  {}
func (m MockLogger) WARN(ctx context.Context, msg string, args ...interface{})  {}
func (m MockLogger) ERROR(ctx context.Context, msg string, args ...interface{}) {}
func (m MockLogger) FATAL(ctx context.Context, msg string, args ...interface{}) {}
func (m MockLogger) Log(level logging.LogLevel, ctx context.Context, msg string, args ...interface{}) {
}
func (m MockLogger) With(args ...any)                   {}
func (m MockLogger) WithGroup(name string, args ...any) {}
func (m MockLogger) WithContext(ctx context.Context)    {}
func (m MockLogger) Clear()                             {}
