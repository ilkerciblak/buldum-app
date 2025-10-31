package account

import (
	"database/sql"
	"net/http"

	"github.com/ilkerciblak/buldum-app/service/account/internal/application"
	repo "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/sql_repository"
	sqlrepository "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/repository/sql_repository"
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

	contactInformationRepository := sqlrepository.NewSqlContactInformationRepository(*accountQueries)
	contactInformationService := application.NewContactInformationService(
		contactInformationRepository,
		accountRepository,
		logger,
	)

	// Contact Information Things
	createContactInformationEP := presentation.NewContactInformationCreateEndPoint(contactInformationService)
	updateContactInformationEP := presentation.NewContactInformationUpdateEndPoint(contactInformationService)
	archiveContactInformationEP := presentation.NewContactInformationArchiveEndPoint(contactInformationService)
	getAllContactInformationByUserEP := presentation.NewContactInformationGetAllByUserEndPoint(contactInformationService)
	getAllContactInformationEP := presentation.NewContactInformationGetAllEndPoint(contactInformationService)
	// Register Endpoints to Mux
	accountMux.HandleFunc(
		createContactInformationEP.Path(),
		middleware.ChainMiddlewareWithEndPoint(
			corepresentation.GenerateHandlerFuncFromEndPoint(
				createContactInformationEP,
				logger,
			),
		),
	)

	accountMux.HandleFunc(
		updateContactInformationEP.Path(),
		middleware.ChainMiddlewareWithEndPoint(
			corepresentation.GenerateHandlerFuncFromEndPoint(
				updateContactInformationEP,
				logger,
			),
		),
	)

	accountMux.HandleFunc(
		archiveContactInformationEP.Path(),
		middleware.ChainMiddlewareWithEndPoint(
			corepresentation.GenerateHandlerFuncFromEndPoint(
				archiveContactInformationEP,
				logger,
			),
		),
	)

	accountMux.HandleFunc(
		getAllContactInformationByUserEP.Path(),
		middleware.ChainMiddlewareWithEndPoint(
			corepresentation.GenerateHandlerFuncFromEndPoint(
				getAllContactInformationByUserEP,
				logger,
			),
		),
	)

	accountMux.HandleFunc(
		getAllContactInformationEP.Path(),
		middleware.ChainMiddlewareWithEndPoint(
			corepresentation.GenerateHandlerFuncFromEndPoint(
				getAllContactInformationEP,
				logger,
			),
		),
	)

	return accountMux

}
