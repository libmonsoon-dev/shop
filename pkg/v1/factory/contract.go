package factory

import "context"

type Interface[T any] interface {
	Create() (T, error)
}

type ContextInterface[T any] interface {
	CreateContext(ctx context.Context) (T, error)
}
