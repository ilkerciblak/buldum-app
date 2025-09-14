package middleware

import (
	"log"
	"net/http"

	"github.com/ilkerciblak/buldum-app/shared/core/domain"
	"github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type PanicRecoverMiddleware struct {
	// TODO LOGGER
}

func (p PanicRecoverMiddleware) Act(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic Recovered From %v", err)

				presentation.RespondWithProblemDetails(w, &domain.InternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	}
}
