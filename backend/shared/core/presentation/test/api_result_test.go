package test

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/ilkerciblak/buldum-app/shared/core/domain"
	"github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

func TestCorePresentation__RespondWithJSON(t *testing.T) {

	cases := []struct {
		Name           string
		Input          any
		ExpectedOutput struct {
			StatusCode int
			Body       []byte
		}
		DoesExpectErrror bool
	}{
		{
			Name:             "Status 200, with foo:bar key values",
			DoesExpectErrror: false,
			Input:            1,
			ExpectedOutput: struct {
				StatusCode int
				Body       []byte
			}{
				StatusCode: 200,
				Body:       []byte(`1`),
			},
		},
		{
			Name:             "Successfull operation with nil payload should give 200 StatusCode and no payload",
			DoesExpectErrror: false,
			Input:            nil,
			ExpectedOutput: struct {
				StatusCode int
				Body       []byte
			}{
				StatusCode: 200,
				Body:       nil,
			},
		},
	}

	for _, tc := range cases {
		t.Run(
			tc.Name,
			func(t *testing.T) {
				testReader := httptest.NewRecorder()

				presentation.RespondWithJSON(testReader, tc.Input)

				if testReader.Result().StatusCode != tc.ExpectedOutput.StatusCode {
					t.Fatalf("Expected StatusCode %v, Output StatusCode %v", testReader.Result().StatusCode, tc.ExpectedOutput.StatusCode)
				}

				if !bytes.Equal(testReader.Body.Bytes(), tc.ExpectedOutput.Body) {
					t.Fatalf("Expected Body %v Got %v", tc.ExpectedOutput.Body, testReader.Body.Bytes())

				}

			},
		)
	}
}

func TestPresentation__RespondWithErrorJson(t *testing.T) {

	validationexp := domain.ValidationException
	validationexp.Errors = map[string]string{
		"name": "name field is required",
	}

	cases := []struct {
		Name            string
		DoesExpectError bool
		Input           domain.ApplicationException
		ExpectedOutput  struct {
			StatusCode int
			Payload    []byte
		}
	}{
		{
			Name:            "Validation Error With Multiple Key:Value Pairs with StatusCode 422",
			DoesExpectError: true,
			Input:           validationexp,
			ExpectedOutput: struct {
				StatusCode int
				Payload    []byte
			}{
				StatusCode: 422,
				Payload:    []byte(`{"code":422,"title":"Unprocessable Entity","errors":{"name":"name field is required"}}`),
			},
		},
		{
			Name:            "Error with 401 and Unathorized Message || Or Any Code & Message Paired Exception",
			DoesExpectError: false,
			Input:           domain.UserNotAuthenticated,
			ExpectedOutput: struct {
				StatusCode int
				Payload    []byte
			}{
				StatusCode: 401,
				Payload:    []byte(`{"code":401,"title":"User Not Authenticated"}`),
			},
		},
	}

	for _, tc := range cases {
		t.Run(
			tc.Name,
			func(t *testing.T) {
				testReader := httptest.NewRecorder()

				presentation.RespondWithErrorJson(testReader, &tc.Input)

				if testReader.Result().StatusCode != tc.ExpectedOutput.StatusCode {

					t.Fatalf("Expected StatusCode %v, Got %v", tc.ExpectedOutput.StatusCode, testReader.Result().StatusCode)
				}

				if !bytes.Equal(testReader.Body.Bytes(), tc.ExpectedOutput.Payload) {
					t.Log(testReader.Body)
					t.Fatalf("Expected Payload %v, Got %v", tc.ExpectedOutput.Payload, testReader.Body.Bytes())
				}
			},
		)
	}
}
