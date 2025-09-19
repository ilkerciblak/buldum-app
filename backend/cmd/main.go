package main

import (
	apiserver "github.com/ilkerciblak/buldum-app/api/api_server"
	appconfig "github.com/ilkerciblak/buldum-app/api/config"
	dbconn "github.com/ilkerciblak/buldum-app/api/db_conn"
)

func main() {

	appConfig, _ := appconfig.NewAppConfig()

	conn, err := dbconn.InitializeSQLDBConnection(appConfig.DB_DRIVER, appConfig.DB_URL)
	if err != nil {
		panic(err)
	}

	apiserver := apiserver.NewApiServer(appConfig, conn)

	apiserver.ListenAndServeApiServer()

}
