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
	"time"

	appconfig "github.com/ilkerciblak/buldum-app/api/config"
	"github.com/ilkerciblak/buldum-app/api/middleware"
	"github.com/ilkerciblak/buldum-app/service/account"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	"github.com/ilkerciblak/buldum-app/shared/core/presentation"
	_ "github.com/lib/pq"
)

type ApiServer struct {
	ServerAddr string
	DB         struct {
		DriverName       string
		ConnectionString string
	}
}

func NewApiServer(cfg *appconfig.AppConfig) *ApiServer {
	return &ApiServer{
		ServerAddr: fmt.Sprintf(":%v", cfg.PORT),
		DB: struct {
			DriverName       string
			ConnectionString string
		}{
			DriverName:       cfg.DB_DRIVER,
			ConnectionString: cfg.DB_URL,
		},
	}
}

func (a *ApiServer) ListenAndServeApiServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	wg := &sync.WaitGroup{}

	server, errchan := a.startHttpServer(wg)

	_ = a.gracefullShutdown(ctx, server, errchan)

	cancel()

	wg.Wait()
}

func (a *ApiServer) gracefullShutdown(ctx context.Context, server *http.Server, errorChan chan error) error {

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	select {
	case err := <-errorChan:
		log.Fatalf("[ERROR]: Starting HTTP Server Failed due: %v", err)
		return err
	case sig := <-signalChan:
		log.Printf("HTTP Server is shutting down gracefully due: %v signal", sig)

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("[ERROR]: Gracefull Shutdown Failed due: %v", err)
			server.Close()
			return err
		}

	}

	return nil

}

func (a *ApiServer) startHttpServer(wg *sync.WaitGroup) (*http.Server, chan error) {

	wg.Add(1)
	defer wg.Done() // samme with wg.Add(-1)
	errChan := make(chan error, 1)

	mux := http.NewServeMux()

	if err := a.registerHandlers(mux); err != nil {
		panic(err)
	}

	conn, err := a.InitializeSQLDBConnection()
	if err != nil {
		errChan <- err
	}

	a.registerDomains(mux, conn)

	server := http.Server{
		Addr:    a.ServerAddr,
		Handler: mux,
	}

	go func() {
		log.Printf("----------------------------------------------")
		log.Printf("---------------Starting Buldum App HTTP Server---------------")
		log.Printf("---------------Listening: %v                  ---------------", a.ServerAddr)
		log.Printf("----------------------------------------------")
		defer conn.Close()
		errChan <- server.ListenAndServe()
	}()

	return &server, errChan

}

func (a *ApiServer) InitializeSQLDBConnection() (*sql.DB, error) {
	db, err := sql.Open(a.DB.DriverName, a.DB.ConnectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil

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

func (a *ApiServer) registerDomains(mux *http.ServeMux, db *sql.DB) {
	account.RegisterAccountDomain(mux, db)

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
