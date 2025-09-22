package main

import (
	"context"
	"sync"
	"time"

	apiserver "github.com/ilkerciblak/buldum-app/api/api_server"
	appconfig "github.com/ilkerciblak/buldum-app/api/config"
	dbconn "github.com/ilkerciblak/buldum-app/api/db_conn"
	"github.com/ilkerciblak/buldum-app/service/account"
)

func main() {

	var wg sync.WaitGroup
	defer wg.Wait()
	errChan := make(chan error, 1)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Initializing the Application Configurations
	appConfig, _ := appconfig.NewAppConfig()
	// Initializing the DBConnection Configurations
	dbConfig := dbconn.NewSqlConnectionConfig(appConfig.DB_DRIVER, appConfig.DB_URL)

	// Initializing Db Connection
	conn := dbConfig.InitializeSQLDBConnection(errChan)
	defer conn.Close()

	apiserver := apiserver.NewApiServer(appConfig, conn)
	apiserver.ConfigureHTTPServer(account.RegisterAccountDomainAPI)
	apiserver.StartHttpServer(errChan, &wg)
	apiserver.GracefullShutdown(ctx, errChan)

}
