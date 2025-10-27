package account

import (
	"database/sql"
	"net/http"

	"github.com/ilkerciblak/buldum-app/service/account/internal/application"
	repo "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/sql_repository"
	account_db "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql"
	presentation "github.com/ilkerciblak/buldum-app/service/account/internal/presentation"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
	"github.com/ilkerciblak/buldum-app/shared/logging"
	"github.com/ilkerciblak/buldum-app/shared/middleware"
)

func RegisterAccountDomainAPI(db *sql.DB, logger logging.ILogger, authMiddleware middleware.IMiddleware) *http.ServeMux {

	accountMux := http.NewServeMux()
	accountQueries := account_db.New(db)
	accountRepository := repo.NewSqlAccountRepository(*accountQueries)
	accountService := application.AccountService(accountRepository, logger)

	createAccountEndPoint := presentation.CreateAccountEndPoint{
		Service: accountService,
		Logger:  logger,
	}
	getAllAccountsEndPoint := presentation.GetAllProfilesEndPoint{
		Service: accountService,
	}
	getAccountByIdEndPoint := presentation.AccountGetByIdEndPoint{
		Service: accountService,
	}
	archiveAccountEndPoint := presentation.ArchiveAccountEndPoint{
		Service: accountService,
	}

	updateAccountEndPoint := presentation.UpdateAccountEndPoint{
		Service: accountService,
		Logger:  logger,
	}

	accountMux.HandleFunc(
		createAccountEndPoint.Path(),
		middleware.ChainMiddlewareWithEndPoint(
			corepresentation.GenerateHandlerFuncFromEndPoint(createAccountEndPoint, logger),
		),
	)

	accountMux.HandleFunc(
		getAccountByIdEndPoint.Path(),

		middleware.ChainMiddlewareWithEndPoint(
			corepresentation.GenerateHandlerFuncFromEndPoint(getAccountByIdEndPoint, logger),
			authMiddleware,
		),
	)

	accountMux.HandleFunc(
		getAllAccountsEndPoint.Path(),
		middleware.ChainMiddlewareWithEndPoint(
			corepresentation.GenerateHandlerFuncFromEndPoint(getAllAccountsEndPoint, logger),
		),
	)

	accountMux.HandleFunc(
		archiveAccountEndPoint.Path(),
		middleware.ChainMiddlewareWithEndPoint(
			corepresentation.GenerateHandlerFuncFromEndPoint(archiveAccountEndPoint, logger),
		),
	)
	accountMux.HandleFunc(
		updateAccountEndPoint.Path(),
		middleware.ChainMiddlewareWithEndPoint(
			corepresentation.GenerateHandlerFuncFromEndPoint(updateAccountEndPoint, logger),
		),
	)
	return accountMux

}
