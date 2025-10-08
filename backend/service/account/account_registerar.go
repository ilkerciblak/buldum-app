package account

import (
	"database/sql"
	"net/http"

	repo "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/sql_repository"
	account_db "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql"
	presentation "github.com/ilkerciblak/buldum-app/service/account/internal/presentation/profile"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
	"github.com/ilkerciblak/buldum-app/shared/middleware"
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
	archiveAccountEndPoint := presentation.ArchiveAccountEndPoint{
		Repository: accountRepository,
	}

	updateAccountEndPoint := presentation.UpdateAccountEndPoint{
		Repository: accountRepository,
	}

	accountMux.HandleFunc(
		createAccountEndPoint.Path(),
		middleware.ChainMiddlewareWithEndPoint(
			corepresentation.GenerateHandlerFuncFromEndPoint(createAccountEndPoint),
		),
	)

	accountMux.HandleFunc(
		getAccountByIdEndPoint.Path(),

		middleware.ChainMiddlewareWithEndPoint(
			corepresentation.GenerateHandlerFuncFromEndPoint(getAccountByIdEndPoint),
		),
	)

	accountMux.HandleFunc(
		getAllAccountsEndPoint.Path(),
		middleware.ChainMiddlewareWithEndPoint(
			corepresentation.GenerateHandlerFuncFromEndPoint(getAllAccountsEndPoint),
		),
	)

	accountMux.HandleFunc(
		archiveAccountEndPoint.Path(),
		middleware.ChainMiddlewareWithEndPoint(
			corepresentation.GenerateHandlerFuncFromEndPoint(archiveAccountEndPoint),
		),
	)
	accountMux.HandleFunc(
		updateAccountEndPoint.Path(),
		middleware.ChainMiddlewareWithEndPoint(
			corepresentation.GenerateHandlerFuncFromEndPoint(updateAccountEndPoint),
		),
	)
	return accountMux

}
