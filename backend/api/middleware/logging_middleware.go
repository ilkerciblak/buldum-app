package middleware

import (
	"log"
	"net/http"
)

type LoggingMiddleware struct {
}

func (l LoggingMiddleware) Act(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logging Before next.ServeHTTP")
		handler.ServeHTTP(w, r)
		log.Printf("Logging After next.ServeHTTP")
	}
}
