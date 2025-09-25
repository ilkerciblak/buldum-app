package command_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/command"
	"github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/mock"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

func TestAccountCommand__ArchiveAccountCommandHandler(t *testing.T) {
	cases := []struct {
		Name            string
		Input           command.ArchiveAccountCommand
		DoesExpectError bool
		ExpectedError   coredomain.IApplicationError
	}{
		{
			Name: "User Already Archived Must Return Error",
			Input: command.ArchiveAccountCommand{
				Id: uuid.Max,
			},
			DoesExpectError: true,
			ExpectedError:   coredomain.InternalServerError,
		},
		{
			Name: "Should OK",
			Input: command.ArchiveAccountCommand{
				Id: uuid.New(),
			},
			DoesExpectError: false,
			ExpectedError:   nil,
		},
	}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {
				err := c.Input.Handler(&mock.MockAccountRepository{}, context.Background())
				if err != nil && !c.DoesExpectError {
					t.Fatalf("Unexpected Error Occurred: %v", err)
				}

				if c.DoesExpectError && (err.GetCode() != c.ExpectedError.GetCode()) {
					t.Fatalf("Error Expectations Not Satisfied\n Expected %v\n Got %v", c.ExpectedError, err)
				}
			},
		)
	}
}
