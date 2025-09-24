package command_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/ilkerciblak/buldum-app/service/account/internal/application/command"
	"github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/mock"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

func Test__CreateAccountCommand_Validation(t *testing.T) {
	cases := []struct {
		Name            string
		Input           command.CreateAccountCommand
		ExpectedOutput  command.CreateAccountCommand
		ExpectedError   coredomain.IApplicationError
		DoesExpectError bool
	}{
		{
			Name: "Invalid Command Request Must Return Validation Error",
			Input: command.CreateAccountCommand{
				Username:  "",
				AvatarUrl: "url",
			},
			DoesExpectError: true,
			ExpectedError: coredomain.RequestValidationError.WithErrors(map[string]string{
				"username": "Username field is required",
			}),
		},
	}

	for _, tc := range cases {
		t.Run(
			tc.Name,
			func(t *testing.T) {
				output, err := tc.Input.Validate()
				if tc.DoesExpectError != (err != nil) {
					t.Fatalf("Test Expecting to Error %v, But err is %v", tc.DoesExpectError, err)
				}

				if tc.DoesExpectError && (err != nil) && !reflect.DeepEqual(tc.ExpectedError, err) {
					t.Fatalf("Test does not satisfies error types, Got %v, Expected %v", err, tc.ExpectedError)
				}

				if !tc.DoesExpectError && !reflect.DeepEqual(*output, tc.ExpectedOutput) {
					t.Fatalf("Test does not satisfies, Got %v, Expected %v", output, tc.ExpectedOutput)
				}

			},
		)
	}

}
func Test__CreateAccountCommand__Handler(t *testing.T) {

	cases := []struct {
		Name            string
		Input           command.CreateAccountCommand
		DoesExpectError bool
	}{
		{
			Name: "Valid Command Should Pass",
			Input: command.CreateAccountCommand{
				Username:  "ilkerciblak",
				AvatarUrl: "url",
			},
			DoesExpectError: false,
		},
		{
			Name: "In-Valid Command Should Raise an Error",
			Input: command.CreateAccountCommand{
				AvatarUrl: "url",
			},
			DoesExpectError: true,
		},
	}

	for _, c := range cases {

		if err := c.Input.Handler(&mock.MockAccountRepository{}, context.Background()); (err != nil) != c.DoesExpectError {
			t.Fatalf("Error Expectations Failed Got %v DoesExpected %v", err, c.DoesExpectError)
		}
	}

}
