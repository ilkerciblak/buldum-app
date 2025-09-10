package apiserver

import "net/http"

type ApiServer struct {
	ServerAddr string
}

func NewApiServer() *ApiServer {
	return &ApiServer{
		ServerAddr: ":8000",
	}
}

func (a *ApiServer) StartServer() {
	mux := http.NewServeMux()

	a.registerHandlers(mux)

	server := http.Server{
		Addr:    a.ServerAddr,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func (a *ApiServer) registerHandlers(mux *http.ServeMux) error {

	mux.HandleFunc("GET /",
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Header().Add("Content-Type", "application/json")
			if _, err := w.Write([]byte(`{"message" : "IT IS ALIVE!}`)); err != nil {
				panic(err)
			}
		})

	return nil
}
