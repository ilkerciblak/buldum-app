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
	handlerFunc := PanicRecoverMiddleware{}.Act(JsonResponseHandlerMiddleware{}.Act(endPoint))
	// var handlerFunc http.HandlerFunc = end
	for _, middleware := range middlewares {
		handlerFunc = middleware.Act(handlerFunc)
	}
	return handlerFunc
}

// CreateMiddlewareChain function that accepts optional list of middlewares implements IMiddleware interface and creates
// pre-defined middleware chains with returning
//
//	func(endPoint presentation.IEndPoint, middlewares ...IMiddleware) http.HandlerFunc
//
// Example Usage:
//
//	//Creating Authenticated Middleware Chain with Logging Middleware
//	authenticatedChain := CreateMiddlewareChain(&AuthenticationMiddleware{}, &LoggingMiddleware{});
//
//	//Let mux handle this path with another middlewares
//	mux.Handle(&ExampleEndPoint{}.Path(), authenticatedChain(&ExampleEndPoint{}, &AnotherMiddleware{}, &Another2Middleware{}))
//
// -	In Default behavior, CreateMiddlewareChain handles http responses with JsonResponseHandlerMiddleware{}
func CreateMiddlewareChain(chainMiddlewares ...IMiddleware) func(endPoint presentation.IEndPoint, middlewares ...IMiddleware) http.HandlerFunc {
	return func(endPoint presentation.IEndPoint, middlewares ...IMiddleware) http.HandlerFunc {
		chainMiddlewares = append(chainMiddlewares, middlewares...)
		handlerFunc := JsonResponseHandlerMiddleware{}.Act(endPoint)
		for _, middleware := range chainMiddlewares {
			handlerFunc = middleware.Act(handlerFunc)
		}
		return handlerFunc
	}
}
