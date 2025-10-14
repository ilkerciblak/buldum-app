package middleware

import (
	"net/http"
)

type IMiddleware interface {
	Act(handlerFunc http.HandlerFunc) http.HandlerFunc
}

func CreateMiddlewareChain(middlewares ...IMiddleware) func(http.Handler) http.Handler {
	return func(hf http.Handler) http.Handler {
		for _, middleware := range middlewares {
			hf = middleware.Act(hf.ServeHTTP)
		}
		return hf
	}
}

func ChainMiddlewareWithEndPoint(endPointHandler http.HandlerFunc, middlewares ...IMiddleware) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, middleware := range middlewares {
			endPointHandler = middleware.Act(endPointHandler)
		}

		endPointHandler.ServeHTTP(w, r)
	}
}
