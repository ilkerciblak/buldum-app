package presentation_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/mock"
	presentation "github.com/ilkerciblak/buldum-app/service/account/internal/presentation/profile"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

func TestEndPoint__ArchiveAccountEndPoint(t *testing.T) {
	archiveAccountEndPoint := presentation.ArchiveAccountEndPoint{
		Repository: &mock.MockAccountRepository{},
	}

	mux := http.NewServeMux()
	mux.HandleFunc(archiveAccountEndPoint.Path(), corepresentation.GenerateHandlerFuncFromEndPoint(archiveAccountEndPoint))

	cases := []struct {
		Name            string
		TestRequest     *http.Request
		DoesExpectError bool
		ExpectedError   coredomain.IApplicationError
	}{

		{
			Name:            "Given Uuid Should Result Internal Server Error",
			TestRequest:     httptest.NewRequest("POST", fmt.Sprintf("/accounts/%s/archive", uuid.Max), nil),
			DoesExpectError: true,
			ExpectedError:   coredomain.InternalServerError,
		},
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
			TestRequest:     httptest.NewRequest("POST", fmt.Sprintf("/accounts/%s/archive", uuid.New()), nil),
			DoesExpectError: false,
			ExpectedError:   nil,
		},
	}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {
				testResponseWriter := httptest.NewRecorder()
				mux.ServeHTTP(testResponseWriter, c.TestRequest)
				a, err := archiveAccountEndPoint.HandleRequest(testResponseWriter, c.TestRequest)
				if c.DoesExpectError {
					if err == nil {
						t.Fatalf("Error Expectations was not satisfied")
					}

					if err.GetCode() != c.ExpectedError.GetCode() {
						t.Fatalf("Error Expectations was not satisfied\nGot %v\n Expected %v", err, c.ExpectedError)
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
