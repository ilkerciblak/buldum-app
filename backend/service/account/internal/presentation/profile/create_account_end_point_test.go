package presentation_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/mock"
	presentation "github.com/ilkerciblak/buldum-app/service/account/internal/presentation/profile"
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
			TestRequest:     httptest.NewRequest("POST", "/accounts", bytes.NewReader([]byte(`{"user_name":"ilkerciblak"}`))),
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
			TestRequest:     httptest.NewRequest("POST", "/accounts", bytes.NewReader([]byte(`{"username":"ilkerciblak}`))),
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
					Context:        ctx,
				}

				res := createAccountEndPoint.HandleRequest(testResponseRecorder.ResponseWriter, c.TestRequest)
				if c.DoesExpectError {
					if res.Error == nil {
						t.Fatalf("Error Expectations was not satisfied")
					}

					if res.Error.(coredomain.IApplicationError).GetCode() != c.ExpectedError.GetCode() {
						t.Fatalf("Error Expectations was not satisfied\nGot %v,Expected %v", res.Error, c.ExpectedError)
					}

				} else {
					if res.Error != nil {
						t.Fatalf("Test was not expecting error but\nGot %v\nExpected %v\n", res.Error, http.StatusCreated)
					}

					if res.Data != nil {
						t.Fatalf("Response Data is not as expected, Expected no data, Got %v", res.Data)
					}

					if res.StatusCode != http.StatusCreated {
						t.Fatalf("Response Status Code is not as expected, Expects %d, Got %d", http.StatusCreated, res.StatusCode)
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
