package account

import (
	"database/sql"
	"net/http"

	"github.com/ilkerciblak/buldum-app/api/middleware"
	repo "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/sql_repository"
	account_db "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql"
	presentation "github.com/ilkerciblak/buldum-app/service/account/internal/presentation"
)

func RegisterAccountDomainAPI(db *sql.DB) *http.ServeMux {

	accountMux := http.NewServeMux()
	accountQueries := account_db.New(db)
	accountRepository := repo.NewSqlAccountRepository(*accountQueries)

	createAccountEndPoint := presentation.CreateAccountEndPoint{
		Repository: accountRepository,
	}
	getAllAccountsEndPoint := presentation.GetAllProfilesEndPoint{
		Repository: accountRepository,
	}
	getAccountByIdEndPoint := presentation.AccountGetByIdEndPoint{
		Repository: accountRepository,
	}

	accountMux.HandleFunc(
		createAccountEndPoint.Path(),
		middleware.ChainMiddlewaresWithEndpoint(
			createAccountEndPoint,
			middleware.LoggingMiddleware{},
		),
	)

	accountMux.HandleFunc(
		getAccountByIdEndPoint.Path(),
		middleware.ChainMiddlewaresWithEndpoint(
			getAccountByIdEndPoint,
			middleware.LoggingMiddleware{},
		),
	)
	accountMux.HandleFunc(
		getAllAccountsEndPoint.Path(),
		middleware.ChainMiddlewaresWithEndpoint(
			getAllAccountsEndPoint,
			middleware.LoggingMiddleware{},
		),
	)

	return accountMux

}
