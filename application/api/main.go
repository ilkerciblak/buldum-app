package main

import (
	apiserver "github.com/ilkerciblak/buldum-app/api/api_server"
	appconfig "github.com/ilkerciblak/buldum-app/shared/config"
)

func main() {

	appConfig, _ := appconfig.NewAppConfig()

	apiserver := apiserver.NewApiServer(appConfig)

	apiserver.ListenAndServeApiServer()

}
