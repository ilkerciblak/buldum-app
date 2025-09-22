package query_test

import (
	"context"
	"testing"

	"github.com/ilkerciblak/buldum-app/service/account/internal/application/query"
	"github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/mock"
)

func Test_GetAll_Query(t *testing.T) {
	mockRepository := &mock.MockAccountRepository{}
	query := &query.AccountGetAllQuery{}
	data, err := query.Handler(mockRepository, context.Background())
	if err != nil {
		t.Fatalf("Error Occured: %v", err)
	}

	if len(data) != 2 {
		t.Fatalf("Data was not like expected")
	}
}
