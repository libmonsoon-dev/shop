//go:generate mockgen -source=./${GOFILE} -destination=../mock/${GOPACKAGE}/mock.go

package repository

import (
	"context"

	shopapi "github.com/libmonsoon-dev/shop/pkg/v1/api"
)

type ItemRepository interface {
	// Creator[shopapi.Item]
	Create(context.Context, *shopapi.Item) error
}

type AttributeRepository interface {
	// Creator[shopapi.Attribute]
	Create(context.Context, *shopapi.Attribute) error
}

type ValueTypeRepository interface {
	// Saver[shopapi.ValueType]
	Save(context.Context, *shopapi.ValueType) error
}
