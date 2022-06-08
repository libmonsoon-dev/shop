package repository

import (
	"context"

	shopapi "github.com/libmonsoon-dev/shop/pkg/v1/api"
)

type Creator[T Entity] interface {
	Create(context.Context, *T) error
}

type Reader[T Entity] interface {
	Read(context.Context, *T) error
}

type Updater[T Entity] interface {
	Update(context.Context, *T) error
}

type Deleter[T Entity] interface {
	Delete(context.Context, shopapi.ID) error
}

// Saver is the interface that wraps the basic Save method.
//
// This method create new record in percipient storage and shouldn't return error if record already exist
// Save should update given Entity with data that is received as a result of saving (such as ID)
type Saver[T Entity] interface {
	Save(context.Context, *T) error
}

type CRUD[T Entity] interface {
	Creator[T]
	Reader[T]
	Updater[T]
	Deleter[T]
}
