package presentation_test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/mock"
	presentation "github.com/ilkerciblak/buldum-app/service/account/internal/presentation/profile"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
	"github.com/ilkerciblak/buldum-app/shared/logging"
)

type MockLogger struct {
	log.Logger
}

func (m *MockLogger) DEBUG(ctx context.Context, msg string, args ...interface{}) {}
func (m *MockLogger) INFO(ctx context.Context, msg string, args ...interface{})  {}
func (m *MockLogger) WARN(ctx context.Context, msg string, args ...interface{})  {}
func (m *MockLogger) ERROR(ctx context.Context, msg string, args ...interface{}) {}
func (m *MockLogger) FATAL(ctx context.Context, msg string, args ...interface{}) {}
func (m *MockLogger) Log(level logging.LogLevel, ctx context.Context, msg string, args ...interface{}) {
}
func (m *MockLogger) With(args ...any)                   {}
func (m *MockLogger) WithGroup(name string, args ...any) {}
func (m *MockLogger) WithContext(ctx context.Context)    {}
func (m *MockLogger) Clear()                             {}

func TestEndPoint__ArchiveAccountEndPoint(t *testing.T) {
	archiveAccountEndPoint := presentation.ArchiveAccountEndPoint{
		Repository: &mock.MockAccountRepository{},
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
