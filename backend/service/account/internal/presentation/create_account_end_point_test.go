package presentation_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/mock"
	"github.com/ilkerciblak/buldum-app/service/account/internal/presentation"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

func TestEndPoint__HandleRequest(t *testing.T) {
	createAccountEndPoint := presentation.CreateAccountEndPoint{
		Repository: &mock.MockAccountRepository{},
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
			TestRequest:     httptest.NewRequest("POST", createAccountEndPoint.Path(), bytes.NewReader([]byte(`{"user_name":"ilkerciblak"}`))),
			ExpectedError:   nil,
		},
		{
			Name:            "Create Account With wrong fields, INVALID with 422",
			DoesExpectError: true,
			TestRequest:     httptest.NewRequest("POST", createAccountEndPoint.Path(), bytes.NewReader([]byte(`{"username":"ilkerciblak"}`))),
			ExpectedError:   coredomain.RequestValidationError,
		},
		{
			Name:            "Create Account With wrong fields, INVALID with 400",
			DoesExpectError: true,
			TestRequest:     httptest.NewRequest("POST", createAccountEndPoint.Path(), bytes.NewReader([]byte(`{"username":"ilkerciblak}`))),
			ExpectedError:   coredomain.BadRequest,
		},
	}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {
				testResponseWriter := httptest.NewRecorder()
				testResponseRecorder := TestResponseRecorder{
					ResponseWriter: testResponseWriter,
					// Context:        ctx,
				}

				a, err := createAccountEndPoint.HandleRequest(testResponseRecorder.ResponseWriter, c.TestRequest)
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

					if a.StatusCode != http.StatusCreated {
						t.Fatalf("Response Status Code is not as expected, Expects %d, Got %d", http.StatusCreated, a.StatusCode)
					}
				}

			},
		)

	}
}

type TestResponseRecorder struct {
	http.ResponseWriter
	context.Context
}

func (t TestResponseRecorder) Header() http.Header {

	return t.ResponseWriter.Header()
}

func (t TestResponseRecorder) Write(b []byte) (int, error) {
	return t.ResponseWriter.Write(b)
}

func (t TestResponseRecorder) WriteHeader(statusCode int) {
	t.ResponseWriter.WriteHeader(statusCode)
}

func (t *TestResponseRecorder) WithContext(ctx context.Context) {
	t.Context = ctx
}

var ctx context.Context = context.Background()
