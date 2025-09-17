package apiserver

import (
	"context"
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
)

type ApiServer struct {
	ServerAddr string
}

func NewApiServer(cfg *appconfig.AppConfig) *ApiServer {
	return &ApiServer{
		ServerAddr: fmt.Sprintf(":%v", cfg.PORT),
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

	mux := http.NewServeMux()

	if err := a.registerHandlers(mux); err != nil {
		panic(err)
	}

	server := http.Server{
		Addr:    a.ServerAddr,
		Handler: mux,
	}

	errChan := make(chan error, 1)

	go func() {
		log.Printf("----------------------------------------------")
		log.Printf("---------------Starting Buldum App HTTP Server---------------")
		log.Printf("---------------Listening: %v                  ---------------", a.ServerAddr)
		log.Printf("----------------------------------------------")
		errChan <- server.ListenAndServe()
	}()

	return &server, errChan

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

	account.RegisterAccountDomainEndPoints(mux)

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
