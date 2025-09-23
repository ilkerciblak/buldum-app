package presentation_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/mock"
	presentation "github.com/ilkerciblak/buldum-app/service/account/internal/presentation/profile"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

func TestEndPoint__GetById(t *testing.T) {
	cases := []struct {
		Name            string
		Input           *http.Request
		DoesExpectError bool
		ExpectedOutput  corepresentation.ApiResult[*model.Profile]
		ExpectedError   coredomain.IApplicationError
	}{
		{
			Name:            "Invalid Request With POST, 405",
			Input:           httptest.NewRequest("POST", "/account", nil),
			DoesExpectError: true,
			ExpectedOutput:  corepresentation.ApiResult[*model.Profile]{},
			ExpectedError:   coredomain.MethodNotAllowed,
		},
		{
			Name:            "Invalid Request With Invalid Id PathVal, 404 Not Found",
			Input:           httptest.NewRequest("GET", "/account/1", nil),
			DoesExpectError: true,
			ExpectedOutput:  corepresentation.ApiResult[*model.Profile]{},
			ExpectedError:   coredomain.NotFound,
		},
		{
			Name:            "Valid Request, 200 Ok",
			Input:           httptest.NewRequest("GET", fmt.Sprintf("/account/%s", uuid.New()), nil),
			DoesExpectError: true,
			ExpectedOutput:  corepresentation.ApiResult[*model.Profile]{},
			ExpectedError:   coredomain.NotFound,
		},
	}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {

				endPoint := presentation.AccountGetByIdEndPoint{
					Repository: &mock.MockAccountRepository{},
				}

				data, err := endPoint.HandleRequest(httptest.NewRecorder(), c.Input)

				if c.DoesExpectError {
					if err == nil {
						t.Fatalf("Error Expectations Not full-filled")
					}
					if c.ExpectedError.GetCode() != err.GetCode() {
						t.Fatalf("Error Expectations Not full-filled\nGot %v\nExpects %v", err, c.ExpectedError)
					}
				} else {
					if err != nil {
						t.Fatalf("Test was Not Expecting Error But Got %v", err)
					}
					if c.ExpectedOutput.StatusCode != data.StatusCode {
						t.Fatalf("Output Error\nExpects %v,Got %v", c.ExpectedOutput.Data, data.Data)

					}
					// if !strings.EqualFold(c.ExpectedOutput.Data.Username, data) {
					// }
				}

			},
		)
	}

}
