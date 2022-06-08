//go:generate mockgen -source=./${GOFILE} -destination=../mock/${GOPACKAGE}/mock.go

package repository

import (
	"context"

	shopv1 "github.com/libmonsoon-dev/shop/pkg/v1"
)

type ItemRepository interface {
	// Creator[shopapi.Item]
	Create(context.Context, *shopv1.Item) error
}

type AttributeRepository interface {
	// Creator[shopapi.Attribute]
	Create(context.Context, *shopv1.Attribute) error
}

type ValueTypeRepository interface {
	// Saver[shopapi.ValueType]
	Save(context.Context, *shopv1.ValueType) error
}
