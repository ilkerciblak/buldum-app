package application

import "context"

type IRequest interface {
	Validate() error
}

type IQueryHandler[T any] interface {
	Handler(r IRequest, ctx context.Context) (T, error)
}

type ICommandHandler interface {
	Handler(r IRequest, ctx context.Context) error
}
