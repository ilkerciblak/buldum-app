package middleware

import (
	"net/http"
)

type IMiddleware interface {
	Act(handlerFunc http.HandlerFunc) http.HandlerFunc
}

func CreateMiddlewareChain(middlewares ...IMiddleware) func(http.HandlerFunc) http.HandlerFunc {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			for _, middleware := range middlewares {
				hf = middleware.Act(hf)
			}
			hf.ServeHTTP(w, r)
		}
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
