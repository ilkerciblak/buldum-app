package apiserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type ApiServer struct {
	ServerAddr string
}

// Config vericem
func NewApiServer() *ApiServer {
	return &ApiServer{
		ServerAddr: ":8080",
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

	a.registerHandlers(mux)

	server := http.Server{
		Addr:    a.ServerAddr,
		Handler: mux,
	}

	errChan := make(chan error, 1)

	go func() {
		log.Printf("----------------------------------------------")
		log.Printf("---------------Starting Buldum App HTTP Server---------------")
		log.Printf("---------------Listening:                     ---------------")
		log.Printf("----------------------------------------------")
		errChan <- server.ListenAndServe()
	}()

	return &server, errChan

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
