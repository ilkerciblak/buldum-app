package sqlrepository_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	repo "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/sql_repository"
	account_db "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql"
	"github.com/ilkerciblak/buldum-app/shared/core/application"
)

var userId uuid.UUID

func prepareMockDB() (*sql.DB, sqlmock.Sqlmock, *repo.SqlAccountRepository, error) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	if err != nil {
		return nil, nil, nil, err
	}

	repo := repo.NewSqlAccountRepository(*account_db.New(db))

	return db, mock, repo, nil

}

func TestSQLAccountRepository__GetById(t *testing.T) {
	db, mock, repo, err := prepareMockDB()
	if err != nil {
		t.Fatalf("Error occured when preparing mock DB")
	}
	defer db.Close()
	ctx := context.Background()
	rows := mock.NewRows([]string{"id", "username", "avatar_url", "created_at", "updated_at", "deleted_at", "is_archived"}).AddRow(
		userId, "ilkerciblak", "url", time.Now(), nil, nil, false,
	)
	mock.ExpectQuery(`-- name: GetProfileById :one .*`).WithArgs(sqlmock.AnyArg()).WillReturnRows(rows)

	_, err = repo.GetById(ctx, userId)
	if err != nil {
		t.Fatalf("Error Occured While repo.GetById with :\n%v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Mock Expectations Were Not Met with :\n%v", err)
	}

}

func TestSQLAccountRepository__GetAll(t *testing.T) {
	db, mock, repo, err := prepareMockDB()
	if err != nil {
		t.Fatalf("Error occured when preparing mock DB")
	}
	defer db.Close()
	ctx := context.Background()

	rows := mock.NewRows([]string{"id", "username", "avatar_url", "created_at", "updated_at", "deleted_at", "is_archived"}).AddRow(
		userId, "ilkerciblak", "url", time.Now(), nil, nil, false,
	)

	mock.ExpectQuery(`-- name: GetAllProfile :many .*`).WithArgs("created_at", 30, 0).WillReturnRows(rows)
	params, _ := application.NewCommonQueryParameters(map[string]any{})
	_, err = repo.GetAll(ctx, *params)
	if err != nil {
		t.Fatalf("Error Occured While repo.GetAll with :\n%v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Mock Expectations Were Not Met with :\n%v", err)
	}
}

func TestSQLAccountRepository__Create(t *testing.T) {
	db, mock, repo, err := prepareMockDB()
	ctx := context.Background()
	user := &model.Profile{
		Id:         userId,
		Username:   "ilkerciblak",
		AvatarUrl:  "url",
		CreatedAt:  time.Now(),
		IsArchived: false,
	}
	if err != nil {
		t.Fatalf("Error occured when preparing mock DB")
	}
	defer db.Close()
	mock.ExpectExec(
		`-- name: CreateProfile :exec .*`).WithArgs(userId, "ilkerciblak", "url", sqlmock.AnyArg(), false).WillReturnResult(driver.ResultNoRows)

	if err := repo.Create(ctx, user); err != nil {
		t.Fatalf("Error Occured While repo.Create with :%v, user :%v", err.Error(), user)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err.Error())
	}

}

func TestSQLAccountRepository__Delete(t *testing.T) {
	db, mock, repo, err := prepareMockDB()
	if err != nil {
		t.Fatalf("Error occured when preparing mock DB")
	}
	defer db.Close()
	ctx := context.Background()

	mock.ExpectExec(`-- name: DeleteProfile :exec .*`).WithArgs(sqlmock.AnyArg()).WillReturnResult(driver.ResultNoRows)

	if err := repo.Delete(ctx, userId); err != nil {
		t.Fatalf("Error Occured While repo.Delete with :\n%v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Mock Expectations Were Not Met with :\n%v", err)
	}

}

func TestSQLAccountRepository__Update(t *testing.T) {
	db, mock, repo, err := prepareMockDB()
	if err != nil {
		t.Fatalf("Error occured when preparing mock DB")
	}
	defer db.Close()
	ctx := context.Background()
	user := &model.Profile{
		Id:         userId,
		Username:   "ilkerciblak",
		AvatarUrl:  "url2",
		CreatedAt:  time.Now(),
		IsArchived: false,
	}
	mock.ExpectExec(`-- name: UpdateProfile :exec .*`).WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnResult(driver.ResultNoRows)

	if err := repo.Update(ctx, userId, user); err != nil {
		t.Fatalf("Error Occured While repo.Update with :\n%v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Mock Expectations Were Not Met with :\n%v", err)
	}

}
