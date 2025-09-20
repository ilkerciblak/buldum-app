package apiserver

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	appconfig "github.com/ilkerciblak/buldum-app/api/config"
	"github.com/ilkerciblak/buldum-app/api/middleware"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	"github.com/ilkerciblak/buldum-app/shared/core/presentation"
	_ "github.com/lib/pq"
)

type ApiServer struct {
	ServerAddr   string
	DbConnection *sql.DB
	*http.ServeMux
	*http.Server
}

func NewApiServer(cfg *appconfig.AppConfig, conn *sql.DB) *ApiServer {
	return &ApiServer{
		ServerAddr:   fmt.Sprintf(":%v", cfg.PORT),
		DbConnection: conn,
	}
}

func (a *ApiServer) GracefullShutdown(ctx context.Context, errorChan <-chan error) {

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	select {
	case err := <-errorChan:
		log.Fatalf("[ERROR]: Starting HTTP Server Failed due: %v\n", err)

	case sig := <-signalChan:
		log.Printf("HTTP Server is shutting down gracefully due: %v signal", sig)

		if err := a.Server.Shutdown(ctx); err != nil {
			log.Printf("[ERROR]: Gracefull Shutdown Failed due: %v\n", err)
			a.Server.Close()

		}

	}

}

func (a *ApiServer) ConfigureHTTPServer(domainRegistarars ...func(mux *http.ServeMux, db *sql.DB)) {

	a.ServeMux = http.NewServeMux()

	for _, f := range domainRegistarars {
		f(a.ServeMux, a.DbConnection)
	}

	a.Server = &http.Server{
		Addr:    a.ServerAddr,
		Handler: a.ServeMux,
	}
}

func (a *ApiServer) StartHttpServer(errChan chan<- error, wg *sync.WaitGroup) {
	log.Println("Buldum Application HTTP Server")
	log.Println("Serving on:\t", a.ServerAddr)
	log.Println("============================")
	wg.Add(1)
	go func() {
		errChan <- a.Server.ListenAndServe()
		defer wg.Done()
	}()
}

func (a *ApiServer) registerHandlers(mux *http.ServeMux) error {

	// mux.HandleFunc(
	// 	HealthCheckEndPoint{}.Path(),
	// 	middleware.ChainMiddlewaresWithEndpoint(&HealthCheckEndPoint{}, &middleware.LoggingMiddleware{}),
	// )
	panicChain := middleware.CreateMiddlewareChain(&middleware.PanicRecoverMiddleware{})

	mux.HandleFunc(
		HealthCheckEndPoint{}.Path(),
		panicChain(HealthCheckEndPoint{}, middleware.LoggingMiddleware{}),
	)

	// account.RegisterAccountDomain(mux)

	return nil
}

type HealthCheckEndPoint struct {
}

func (h HealthCheckEndPoint) Path() string {
	return "GET /health"
}

func (h HealthCheckEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) (presentation.ApiResult[any], coredomain.IApplicationError) {
	// panic("PanicRecoverDemo")
	return presentation.ApiResult[any]{Data: nil, StatusCode: 200}, nil
}
