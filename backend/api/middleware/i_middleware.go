package middleware

import (
	"net/http"

	"github.com/ilkerciblak/buldum-app/shared/core/presentation"
)

type IMiddleware interface {
	Act(handlerFunc http.HandlerFunc) http.HandlerFunc
}

// presentation.GenerateHandlerFunc()
func ChainMiddlewaresWithEndpoint(endPoint presentation.IEndPoint, middlewares ...IMiddleware) http.HandlerFunc {
	// handlerFunc := presentation.GenerateHandlerFuncFromEndPoint(endPoint)
	handlerFunc := PanicRecoverMiddleware{}.Act(ErrorHandlerMiddleware{}.Act(endPoint))
	// var handlerFunc http.HandlerFunc = end
	for _, middleware := range middlewares {
		handlerFunc = middleware.Act(handlerFunc)
	}

	return handlerFunc.ServeHTTP
}

// func CreateMiddlewareChain(middlewares ...IMiddleware) http.HandlerFunc {
// 	var handler http.HandlerFunc
// 	for i:=0; i<len(middlewares)-1; i++ {
// 		handler = middlewares[i].Act()
// 	}
// }
