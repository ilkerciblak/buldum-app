package command_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/command"
	"github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/mock"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

func TestAccountCommands__UpdateAccountCommandValidate(t *testing.T) {
	cases := []struct {
		Name            string
		Input           command.UpdateAccountCommand
		DoesExpectError bool
		ExpectedOutput  command.UpdateAccountCommand
		ExpectedError   coredomain.IApplicationError
	}{}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {
				output, err := c.Input.Validate()
				if c.DoesExpectError {
					if err == nil {
						t.Fatalf("Error Expectations Not Satisfied, Expected An %v", c.ExpectedError)
					}

					if err.(coredomain.IApplicationError).GetCode() != c.ExpectedError.GetCode() {
						t.Fatalf("Error Expectations Not Satisfied\n Expected An %v\nGot %v", c.ExpectedError, err)
					}
				} else {
					if !reflect.DeepEqual(c.ExpectedOutput, output) {
						t.Fatalf("Output Not Satisfies Expectations\nExpected %v\nGot %v", c.ExpectedOutput, output)
					}
				}
			},
		)
	}
}

func TestAccountCommands__UpdateAccountCommandHandler(t *testing.T) {
	cases := []struct {
		Name            string
		Input           *http.Request
		DoesExpectError bool
		ExpectedError   coredomain.IApplicationError
	}{
		{
			Name:            "Not Giving Username Input, Should 422",
			DoesExpectError: true,
			ExpectedError:   coredomain.RequestValidationError,
			Input:           httptest.NewRequest("PUT", fmt.Sprintf("/accounts/%s", uuid.New()), bytes.NewReader([]byte(``))),
		},
		{
			Name:            "Only Giving Username Input, Should OK",
			DoesExpectError: false,
			ExpectedError:   nil,
			Input:           httptest.NewRequest("PUT", fmt.Sprintf("/accounts/%s", uuid.New()), bytes.NewReader([]byte(`{"user_name":"foobar"}`))),
		},
	}

	for _, tc := range cases {
		t.Run(
			tc.Name,
			func(t *testing.T) {
				mux := http.NewServeMux()
				mux.HandleFunc("PUT /accounts/{user_id}", func(w http.ResponseWriter, r *http.Request) {

					var updateAccountRequest struct {
						Username  string `json:"user_name"`
						AvatarUrl string `json:"avatar_url"`
					}
					defer r.Body.Close()
					_ = json.NewDecoder(r.Body).Decode(&updateAccountRequest)

					userid := r.PathValue("user_id")

					com := command.UpdateAccountCommand{
						Username:  updateAccountRequest.Username,
						AvatarUrl: updateAccountRequest.AvatarUrl,
					}
					com.SetUserID(userid)

					err := com.Handler(&mock.MockAccountRepository{}, r.Context())

					if tc.DoesExpectError {
						if err == nil {
							t.Fatalf("Error Expectations Not Satisfied, Expected An %v", tc.ExpectedError)
						}

						if err.GetCode() != tc.ExpectedError.GetCode() {
							t.Fatalf("Error Expectations Not Satisfied\n Expected An %v\nGot %v", tc.ExpectedError, err)
						}
					} else {
						if err != nil {
							t.Fatalf("Unexpected Error Occurred %v", err)
						}

					}
				})
				mux.ServeHTTP(httptest.NewRecorder(), tc.Input)

			},
		)
	}
}
