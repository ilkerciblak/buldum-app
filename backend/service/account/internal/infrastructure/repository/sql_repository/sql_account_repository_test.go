package sqlrepository_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/model"
	"github.com/ilkerciblak/buldum-app/service/account/internal/domain/repository"
	repo "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/sql_repository"
	account_db "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql"
	"github.com/ilkerciblak/buldum-app/shared/core/application"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
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
	tt := time.Now()
	// _ = mock.NewRows([]string{"id", "username", "avatar_url", "created_at", "updated_at", "deleted_at", "is_archived"}).AddRow(
	// 	userId, "ilkerciblak", "url", tt, nil, nil, false,
	// )

	cases := []struct {
		Name            string
		ExpectedRows    *sqlmock.Rows
		ExpectedData    []*model.Profile
		Params          *application.CommonQueryParameters
		Filter          *repository.ProfileGetAllQueryFilter
		DoesExpectError bool
		ExpectedError   coredomain.IApplicationError
	}{
		{
			Name: "Zottiri zittiri",
			ExpectedRows: sqlmock.NewRows([]string{"id", "username", "avatar_url", "created_at", "updated_at", "deleted_at", "is_archived"}).AddRow(
				userId, "ilkerciblak", "url", tt, nil, nil, false,
			),
			ExpectedData: []*model.Profile{{
				Id:         userId,
				Username:   "ilkerciblak",
				AvatarUrl:  "url",
				CreatedAt:  tt,
				IsArchived: false,
			}},
			Params: &application.CommonQueryParameters{
				Pagination: application.Pagination{
					Page:   1,
					Limit:  10,
					Offset: 0,
				},
				Sorting: application.Sorting{
					Sort:  "created_at",
					Order: "asc",
				},
			},
			Filter:          repository.DefaultAccountGetAllQueryFilter(),
			DoesExpectError: false,
		},
	}

	for _, c := range cases {
		t.Run(
			c.Name,
			func(t *testing.T) {
				ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
				mock.ExpectQuery(`-- name: GetAllProfile :many .*`).WithArgs(
					c.Params.Limit,
					c.Params.Limit*(c.Params.Page-1),
					c.Filter.Username,
					c.Filter.IsArchived,
				).WillReturnRows(c.ExpectedRows)

				data, err := repo.GetAll(ctx, *c.Params, *c.Filter)

				if (err != nil) && !c.DoesExpectError {
					t.Fatalf("Unexpected Error Occured %v", err)
				}
				t.Log(len(data))

				if !reflect.DeepEqual(err, c.ExpectedError) {
					t.Fatalf("Error Output Was Not Satisfied\n Expected %v\n Got %v", c.ExpectedError, err)
				}

				if !reflect.DeepEqual(data, c.ExpectedData) && !c.DoesExpectError {
					t.Fatalf("Output Expectations Not Satisfied\n Expected %v\n Got %v", c.ExpectedData, data)
				}

			},
		)
	}

	// params, _ := application.NewCommonQueryParameters(map[string]any{})
	// filter, _ := repository.NewAccountGetAllQueryFilter(map[string]any{})

	// _, err = repo.GetAll(ctx, *params, *filter)
	// if err != nil {
	// 	t.Fatalf("Error Occured While repo.GetAll with :\n%v", err)
	// }

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Fatalf("Mock Expectations Were Not Met with :\n%v", err)
	// }
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
