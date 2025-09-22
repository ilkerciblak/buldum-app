package middleware

import (
	"log"
	"net/http"

	"github.com/ilkerciblak/buldum-app/shared/core/coredomain"
	corepresentation "github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type PanicRecoverMiddleware struct {
	// TODO LOGGER
}

func (p PanicRecoverMiddleware) Act(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic Recovered From %v", err)

				corepresentation.RespondWithProblemDetails(w, coredomain.InternalServerError.WithMessage(err))
			}
		}()
		log.Printf("PanicRecoverMiddleware Before ServeHTTP")
		next.ServeHTTP(w, r)
		log.Printf("PanicRecoverMiddleware After ServeHTTP")
	}
}
