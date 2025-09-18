package account

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/api/middleware"
	account_db "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql"
	infrastructure "github.com/ilkerciblak/buldum-app/service/account/internal/infrastructure/sql_repository"
	"github.com/ilkerciblak/buldum-app/service/account/internal/presentation"
)

func RegisterAccountDomainEndPoints(mux *http.ServeMux) {

	accountRepository := infrastructure.NewSqlAccountRepository(account_db.Queries{})

	createAccountEndPoint := presentation.CreateAccountEndPoint{
		Repository: accountRepository,
	}

	mux.HandleFunc(createAccountEndPoint.Path(), middleware.ChainMiddlewaresWithEndpoint(
		createAccountEndPoint,
	))

}
