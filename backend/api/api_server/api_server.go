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
	api_middlewares "github.com/ilkerciblak/buldum-app/api/middleware"
	"github.com/ilkerciblak/buldum-app/shared/authentication"
	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	presentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
	"github.com/ilkerciblak/buldum-app/shared/logging"
	"github.com/ilkerciblak/buldum-app/shared/middleware"
	_ "github.com/lib/pq"
)

type ApiServer struct {
	ServerAddr   string
	DbConnection *sql.DB
	secureKey    string
	*http.ServeMux
	*http.Server
}

func NewApiServer(cfg *appconfig.AppConfig, conn *sql.DB) *ApiServer {
	return &ApiServer{
		ServerAddr:   fmt.Sprintf(":%v", cfg.PORT),
		DbConnection: conn,
		secureKey:    cfg.SecureKey,
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

func (a *ApiServer) ConfigureHTTPServer(domainRegisterars ...func(db *sql.DB, logger logging.ILogger, authMiddleware middleware.IMiddleware) *http.ServeMux) {

	a.ServeMux = http.NewServeMux()
	logger := logging.NewSlogger(logging.LoggerOptions{
		MinLevel:    logging.INFO,
		JsonLogging: true,
	})

	loggingMiddleware := api_middlewares.LoggingMiddleware{
		ILogger: logger,
	}

	authMiddleware := api_middlewares.AuthenticationMiddleware{
		ILogger:   logger,
		JWTHelper: *authentication.NewJWTMaker(a.secureKey),
	}

	apiHandler := middleware.CreateMiddlewareChain(&loggingMiddleware, api_middlewares.PanicRecoverMiddleware{})

	for _, f := range domainRegisterars {
		a.ServeMux.Handle("/api/", apiHandler(http.StripPrefix("/api", f(a.DbConnection, logger, &authMiddleware))))
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

type HealthCheckEndPoint struct {
}

func (h HealthCheckEndPoint) Path() string {
	return "GET /health"
}

func (h HealthCheckEndPoint) HandleRequest(w http.ResponseWriter, r *http.Request) (presentation.ApiResult[any], coredomain.IApplicationError) {
	// panic("PanicRecoverDemo")
	return presentation.ApiResult[any]{Data: nil, StatusCode: 200}, nil
}
