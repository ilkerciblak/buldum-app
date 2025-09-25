package presentation_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/mock"
	presentation "github.com/ilkerciblak/buldum-app/service/account/internal/presentation/profile"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

func TestEndPoint__ArchiveAccountEndPoint(t *testing.T) {
	archiveAccountEndPoint := presentation.CreateAccountEndPoint{
		Repository: &mock.MockAccountRepository{},
	}

	cases := []struct {
		Name            string
		TestRequest     *http.Request
		DoesExpectError bool
		ExpectedError   coredomain.IApplicationError
	}{}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {
				testResponseWriter := httptest.NewRecorder()

				a, err := archiveAccountEndPoint.HandleRequest(testResponseWriter, c.TestRequest)
				if c.DoesExpectError {
					if err == nil {
						t.Fatalf("Error Expectations was not satisfied")
					}

					if err.GetCode() != c.ExpectedError.GetCode() {
						t.Fatalf("Error Expectations was not satisfied\nGot %v,Expected %v", err, c.ExpectedError)
					}

				} else {
					if err != nil {
						t.Fatalf("Test was not expecting error but\nGot %v\nExpected %v\n", err, http.StatusCreated)
					}

					if a.Data != nil {
						t.Fatalf("Response Data is not as expected, Expected no data, Got %v", a.Data)
					}

					if a.StatusCode != http.StatusNoContent {
						t.Fatalf("Response Status Code is not as expected, Expects %d, Got %d", http.StatusCreated, a.StatusCode)
					}
				}
			},
		)
	}
}
