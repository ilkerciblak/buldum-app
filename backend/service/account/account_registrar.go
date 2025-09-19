package account

import (
	"database/sql"
	"net/http"

	"github.com/ilkerciblak/buldum-app/api/middleware"
	account_db "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql"
	infrastructure "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql_repository"
	"github.com/ilkerciblak/buldum-app/service/account/internal/presentation"
)

func RegisterAccountDomain(mux *http.ServeMux, db *sql.DB) {
	accountQueries := account_db.New(db)

	accountRepository := infrastructure.NewSqlAccountRepository(*accountQueries)

	createAccountEndPoint := presentation.CreateAccountEndPoint{
		Repository: accountRepository,
	}

	mux.HandleFunc(
		createAccountEndPoint.Path(),
		middleware.ChainMiddlewaresWithEndpoint(
			createAccountEndPoint,
			middleware.LoggingMiddleware{},
		),
	)

}
