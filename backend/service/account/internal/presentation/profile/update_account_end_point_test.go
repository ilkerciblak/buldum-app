package presentation_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/mock"
	presentation "github.com/ilkerciblak/buldum-app/service/account/internal/presentation/profile"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

func TestEndPoint__UpdateAccount(t *testing.T) {
	var endpoint corepresentation.IEndPoint = presentation.UpdateAccountEndPoint{
		Repository: &mock.MockAccountRepository{},
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
					fmt.Sprintf("/accounts/%v", uuid.New()),
					bytes.NewBuffer([]byte(`{"username":"ilkerciblak"}`)),
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
					fmt.Sprintf("/accounts/%v", uuid.New()),
					nil,
				)
			},
			ExpectedStatusCode: http.StatusBadRequest,
		},
		{
			Name: "Malformed JSON Body",
			TestRequest: func() *http.Request {
				return httptest.NewRequest(http.MethodPut, fmt.Sprintf("/accounts/%v", uuid.New()), bytes.NewReader([]byte(`{invalid-json"`)))
			},

			ExpectedStatusCode: http.StatusBadRequest,
		},
		{
			Name: "422 Validation Error",
			TestRequest: func() *http.Request {
				return httptest.NewRequest(http.MethodPut, fmt.Sprintf("/accounts/%v", uuid.New()), bytes.NewReader([]byte(`{"avatar_url":"blabla"}`)))
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
