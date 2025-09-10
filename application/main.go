package main

import (
	apiserver "github.com/ilkerciblak/buldum-app/api/api_server"
)

func main() {

	apiserver := apiserver.NewApiServer()

	apiserver.StartServer()

}
