package test

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type mockEndpoint struct{}

func (m *mockEndpoint) Path() string {
	return "GET /test"
}

func (m *mockEndpoint) HandleRequest(w http.ResponseWriter, r *http.Request) (any, error) {
	defer r.Body.Close()
	if r.Body != nil {

		return r.Body, nil
	}

	return nil, errors.New("Error")
}

func TestCorePresentation__RespondWithJSON(t *testing.T) {

	cases := []struct {
		Name             string
		Input            *http.Request
		ExpectedOutput   *http.Response
		DoesExpectErrror bool
	}{
		{
			Name:             "Status 200, with foo:bar key values",
			DoesExpectErrror: false,
			Input:            httptest.NewRequest("GET", "/test", bytes.NewReader([]byte(`{"foo":"bar"}`))),
			ExpectedOutput: &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader([]byte(`"1"`))),
			},
		},
	}

	for _, tc := range cases {
		t.Run(
			tc.Name,
			func(t *testing.T) {
				testReader := httptest.NewRecorder()

				presentation.RespondWithJSON(testReader, 1)

				if testReader.Result().StatusCode != tc.ExpectedOutput.StatusCode {
					t.Fatalf("Expected StatusCode %v, Output StatusCode %v", testReader.Result().StatusCode, tc.ExpectedOutput.StatusCode)
				}

				t.Logf("%v", testReader.Result().StatusCode)
				t.Logf("%v", testReader.Result().Status)

			},
		)
	}
}
