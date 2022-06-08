//go:generate mockgen -source=./${GOFILE} -destination=../mock/${GOPACKAGE}/mock.go

package unitofwork

import (
	"context"

	"github.com/libmonsoon-dev/shop/pkg/v1/repository"
)

type Interface interface {
	Complete(err *error)
}

type ItemCreatorFactory interface {
	CreateContext(ctx context.Context) (ItemCreator, error)
	// TODO: factory.ContextInterface[ItemCreator]
}

type ItemCreator interface {
	Interface

	ItemRepository() repository.ItemRepository
	AttributeRepository() repository.AttributeRepository
	ValueTypeRepository() repository.ValueTypeRepository
}
