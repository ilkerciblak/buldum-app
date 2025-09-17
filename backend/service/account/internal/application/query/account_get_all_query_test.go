package query_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/application/query"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
)

type MockAccountRepository struct {
}

func (m MockAccountRepository) GetById(userId uuid.UUID) (*model.Profile, error) {
	return model.NewProfile("ilkerciblak", "url"), nil
}
func (m MockAccountRepository) GetAll() ([]*model.Profile, error) {
	return []*model.Profile{
		model.NewProfile("ilkerciblak", "url"),
		model.NewProfile("ilkerciblak", "url"),
	}, nil
}
func (m MockAccountRepository) Create(p *model.Profile) error {
	return nil
}
func (m MockAccountRepository) Update(userId uuid.UUID, p *model.Profile) error {
	return nil
}
func (m MockAccountRepository) Delete(userId uuid.UUID) error {
	return nil
}

func Test_GetAll_Query(t *testing.T) {
	query := &query.AccountGetAllQuery{}
	data, err := query.Handler(MockAccountRepository{}, context.Background())
	if err != nil {
		t.Fatalf("Error Occured: %v", err)
	}

	if len(data) != 2 {
		t.Fatalf("Data was not like expected")
	}
}
