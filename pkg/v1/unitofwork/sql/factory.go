package sqlunitofwork

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/libmonsoon-dev/shop/pkg/v1/factory"
	"github.com/libmonsoon-dev/shop/pkg/v1/unitofwork"
)

func NewFactory(db *sql.DB, opts *sql.TxOptions) *Factory {
	return &Factory{db: db, opts: opts}
}

type Factory struct {
	db   *sql.DB
	opts *sql.TxOptions
}

var _ factory.ContextInterface[*UnitOfWork] = (*Factory)(nil)

func (f *Factory) CreateContext(ctx context.Context) (*UnitOfWork, error) {
	tx, err := f.db.BeginTx(ctx, f.opts)
	if err != nil {
		return nil, fmt.Errorf("begin sql transaction: %w", err)
	}

	return NewUnitOfWork(tx), nil
}

func NewItemCreatorFactory(db *sql.DB) *ItemCreatorFactory {
	f := NewFactory(
		db,
		&sql.TxOptions{Isolation: sql.LevelRepeatableRead},
	)

	return &ItemCreatorFactory{factory: f}
}

type ItemCreatorFactory struct {
	factory *Factory
}

var _ unitofwork.ItemCreatorFactory = (*ItemCreatorFactory)(nil)

func (f *ItemCreatorFactory) CreateContext(ctx context.Context) (unitofwork.ItemCreator, error) {
	unit, err := f.factory.CreateContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("create base unit of work: %w", err)
	}

	return unit, nil
}

//TODO:
//type GenericFactory[T unitofwork.Interface] struct {
//	factory *Factory
//}
//
//func (f *GenericFactory[T]) CreateContext(ctx context.Context) (T, error) {
//	unit, err := f.factory.CreateContext(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("create base unit of work: %w", err)
//	}
//
//	return unit, nil
//}
