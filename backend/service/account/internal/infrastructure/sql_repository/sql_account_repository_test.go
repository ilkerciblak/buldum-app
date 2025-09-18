package sqlrepository_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	account "github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"

	account_db "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql"
	sqlrepository "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql_repository"
	"github.com/stretchr/testify/mock"
)

// Mock sqlc Queries
type MockQueries struct {
	mock.Mock
}

func (m *MockQueries) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	ret := m.Called(ctx, query, args)
	return ret.Get(0).(sql.Result), ret.Error(1)
}

func (m *MockQueries) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	ret := m.Called(ctx, query)
	return ret.Get(0).(*sql.Stmt), ret.Error(1)
}

func (m *MockQueries) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	ret := m.Called(ctx, query, args)

	return ret.Get(0).(*sql.Rows), ret.Error(1)
}

func (m *MockQueries) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	ret := m.Called(ctx, query, args)
	return ret.Get(0).(*sql.Row)
}

// ------------------- TESTS -------------------

func TestSqlAccountRepository_GetById(t *testing.T) {
	ctx := context.Background()
	mockQueries := new(MockQueries)
	m := account_db.New(mockQueries)
	repo := sqlrepository.NewSqlAccountRepository(*m)

	userID := uuid.New()
	expectedProfile := account_db.AccountProfile{
		ID:        userID,
		UserID:    uuid.New(),
		UserName:  "testuser",
		CreatedAt: time.Now(),
	}

	mockQueries.On("GetProfileByID", ctx, userID).Return(expectedProfile, nil)

	profile, err := repo.GetById(ctx, userID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if profile == nil {
		t.Fatal("expected profile, got nil")
	}
	if profile.Username != expectedProfile.UserName {
		t.Fatalf("expected username %s, got %s", expectedProfile.UserName, profile.Username)
	}

	mockQueries.AssertExpectations(t)
}

func TestSqlAccountRepository_GetAll(t *testing.T) {
	ctx := context.Background()
	mockQueries := new(MockQueries)
	m := account_db.New(mockQueries)
	repo := sqlrepository.NewSqlAccountRepository(*m)

	profiles := []account_db.AccountProfile{
		{ID: uuid.New(), UserID: uuid.New(), UserName: "user1", CreatedAt: time.Now()},
		{ID: uuid.New(), UserID: uuid.New(), UserName: "user2", CreatedAt: time.Now()},
	}

	mockQueries.On("ListProfiles", ctx).Return(profiles, nil)

	result, err := repo.GetAll(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != len(profiles) {
		t.Fatalf("expected %d profiles, got %d", len(profiles), len(result))
	}

	mockQueries.AssertExpectations(t)
}

func TestSqlAccountRepository_Create(t *testing.T) {
	ctx := context.Background()
	mockQueries := new(MockQueries)
	m := account_db.New(mockQueries)
	repo := sqlrepository.NewSqlAccountRepository(*m)

	p := &account.Profile{
		Id:         uuid.New(),
		UserID:     uuid.New(),
		Username:   "newuser",
		CreatedAt:  time.Now(),
		IsArchived: false,
	}

	mockQueries.On("CreateProfile", ctx, mock.Anything).Return(account_db.AccountProfile{
		ID: p.Id, UserID: p.UserID, UserName: p.Username, CreatedAt: p.CreatedAt,
	}, nil)

	err := repo.Create(ctx, p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	mockQueries.AssertExpectations(t)
}

func TestSqlAccountRepository_Update(t *testing.T) {
	ctx := context.Background()
	mockQueries := new(MockQueries)
	m := account_db.New(mockQueries)
	repo := sqlrepository.NewSqlAccountRepository(*m)

	userID := uuid.New()
	p := &account.Profile{
		Id:       userID,
		Username: "updateduser",
	}

	mockQueries.On("UpdateProfile", ctx, mock.Anything).Return(account_db.AccountProfile{
		ID: userID, UserName: p.Username,
	}, nil)

	err := repo.Update(ctx, userID, p)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	mockQueries.AssertExpectations(t)
}

func TestSqlAccountRepository_Delete(t *testing.T) {
	ctx := context.Background()
	mockQueries := new(MockQueries)
	m := account_db.New(mockQueries)
	repo := sqlrepository.NewSqlAccountRepository(*m)

	userID := uuid.New()

	mockQueries.On("DeleteProfile", ctx, userID).Return(nil)

	err := repo.Delete(ctx, userID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	mockQueries.AssertExpectations(t)
}
