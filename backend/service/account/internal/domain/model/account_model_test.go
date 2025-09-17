package model_test

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	domain "github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
)

func TestAccountDomain__TestContactInformationType(t *testing.T) {
	cases := []struct {
		Name            string
		TestInput       string
		ExpectedOutput  any
		DoesExpectError bool
	}{
		{
			Name:            "Invalid Type Must Assert Error",
			TestInput:       "Sms",
			ExpectedOutput:  nil,
			DoesExpectError: true,
		},
		{
			Name:            "Valid Type Should Pass",
			TestInput:       "Email",
			ExpectedOutput:  domain.Email,
			DoesExpectError: false,
		},
	}

	for _, tc := range cases {

		t.Run(
			tc.Name,
			func(t *testing.T) {

				output, err := domain.NewContactInformation(uuid.UUID{}, domain.ContactInformationType(tc.TestInput), true).Type.Validate()

				if tc.DoesExpectError != (err != nil) {
					t.Fatalf("Test Does Expect Error: %v, While err is %v", tc.DoesExpectError, err)
				}

				if !tc.DoesExpectError && (!reflect.DeepEqual(*output, tc.ExpectedOutput)) {
					t.Fatalf("Test Expects: %v, Got %v", tc.ExpectedOutput, output)
				}

			},
		)
	}

}
