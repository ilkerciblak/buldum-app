package middleware

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/shared/core/domain"
)

type IMiddleware interface {
	Act(*http.HandlerFunc) domain.IApplicationError
}

func ChainMiddlewares(handlerFunc http.HandlerFunc, middlewares ...IMiddleware) http.HandlerFunc {
	for _, middleware := range middlewares {
		_ = middleware.Act(&handlerFunc)

	}

	return handlerFunc
}

func a() {
	ChainMiddlewares(func(w http.ResponseWriter, r *http.Request) {})
}
