package query_test

import (
	"context"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/query"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/mock"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
)

func TestQuery__GetByIdAccount(t *testing.T) {
	mockRepo := mock.MockAccountRepository{}

	cases := []struct {
		Name            string
		Input           uuid.UUID
		DoesExpectError bool
		ExpectedError   coredomain.IApplicationError
		ExpectedOutput  *model.Profile
	}{
		{
			Name:            "This query should return 404",
			Input:           uuid.Max,
			DoesExpectError: true,
			ExpectedError:   coredomain.NotFound,
			ExpectedOutput:  nil,
		},
		{
			Name:            "This query should return 404",
			Input:           uuid.Nil,
			DoesExpectError: true,
			ExpectedError:   coredomain.NotFound,
			ExpectedOutput:  nil,
		},
		{
			Name:            "This query should return 200 with some data",
			Input:           uuid.New(),
			DoesExpectError: false,
			ExpectedError:   nil,
			ExpectedOutput: &model.Profile{
				Username: "ilkerciblak",
			},
		},
	}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {
				q := &query.AccountGetByIdQuery{
					Id: c.Input,
				}
				data, err := q.Handler(&mockRepo, context.Background())

				if c.DoesExpectError {
					if err == nil {
						t.Fatalf("Error Expectations Not full-filled")
					}
					if c.ExpectedError.GetCode() != err.GetCode() {
						t.Fatalf("Error Expectations Not full-filled\nGot %v\nExpects %v", err, c.ExpectedError)
					}
				} else {
					if err != nil {
						t.Fatalf("Test was Not Expecting Error But Got %v", err)
					}

					if !strings.EqualFold(c.ExpectedOutput.Username, data.Username) {
						t.Fatalf("Output Error\nExpects %v,Got %v", c.ExpectedOutput.Username, data.Username)
					}
				}

			},
		)
	}

}
