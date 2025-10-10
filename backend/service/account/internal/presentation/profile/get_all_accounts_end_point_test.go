package presentation_test

import (
	"reflect"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/mock"
	presentation "github.com/ilkerciblak/buldum-app/service/account/internal/presentation/profile"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

func TestEndPoint__GetAllProfiles(t *testing.T) {
	// endPoint := presentation.

	cases := []struct {
		Name            string
		TestRequest     *http.Request
		ExpectedOutput  corepresentation.ApiResult[[]*model.Profile]
		DoesExpectError bool
		ExpectedError   coredomain.IApplicationError
	}{
		{
			Name:            "Valid Request, 200 OK with Some Data",
			DoesExpectError: false,
			ExpectedError:   nil,
			TestRequest:     httptest.NewRequest("GET", "/account", nil),
			ExpectedOutput: corepresentation.ApiResult[[]*model.Profile]{
				Data: []*model.Profile{
					model.NewProfile("ilkerciblak", "url"),
					model.NewProfile("ilkerciblak", "url"),
				},
				StatusCode: 200,
			},
		},
		{
			Name:            "InValid Request with POST path, Should 405",
			DoesExpectError: true,
			ExpectedError:   coredomain.MethodNotAllowed,
			TestRequest:     httptest.NewRequest("POST", "/account", nil),
			ExpectedOutput: corepresentation.ApiResult[[]*model.Profile]{
				Data:       nil,
				StatusCode: 0,
			},
		},
	}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {
				repo := &mock.MockAccountRepository{}

				endPoint := presentation.GetAllProfilesEndPoint{
					Repository: repo,
				}

				res := endPoint.HandleRequest(httptest.NewRecorder(), c.TestRequest)
				if c.DoesExpectError {
					if res.Error == nil {
						t.Fatalf("Test Error Expectation was not full-filled")
					}

					if res.Error.(coredomain.IApplicationError).GetCode() != c.ExpectedError.GetCode() {
						t.Fatalf("Test Error Expectation was not full-filled\nGot %v\nExpect %v", res.Error, c.ExpectedError)
					}
				} else {
					if res.Error != nil {
						t.Fatalf("Test Was Not Expecting Error But Got %v", res.Error)
					}

					if reflect.DeepEqual(res.Data, c.ExpectedOutput) {
						t.Fatalf("Output Not Satisfied\nGot%v\nExpects%v\n", res.Data, c.ExpectedOutput.Data)
					}

					if res.StatusCode != c.ExpectedOutput.StatusCode {
						t.Fatalf("Expected Status Code %v, Got %v", c.ExpectedOutput.StatusCode, res.StatusCode)
					}
				}

			},
		)
	}

}
