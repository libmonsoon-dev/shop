package sqlunitofwork

import (
	"database/sql"

	"github.com/libmonsoon-dev/shop/pkg/v1/repository"
	"github.com/libmonsoon-dev/shop/pkg/v1/unitofwork"
)

func NewUnitOfWork(tx *sql.Tx) *UnitOfWork {
	return &UnitOfWork{
		UnitOfWork: unitofwork.NewUnitOfWork(tx),
		tx:         tx,
	}
}

var (
	_ unitofwork.Interface   = (*UnitOfWork)(nil)
	_ unitofwork.ItemCreator = (*UnitOfWork)(nil)
)

type UnitOfWork struct {
	unitofwork.UnitOfWork
	tx *sql.Tx

	itemRepositoryFactory      RepositoryFactory[repository.ItemRepository]
	attributeRepositoryFactory RepositoryFactory[repository.AttributeRepository]
	valueTypeRepositoryFactory RepositoryFactory[repository.ValueTypeRepository]
}

func (u *UnitOfWork) ItemRepository() repository.ItemRepository {
	return u.itemRepositoryFactory.Create(u.tx)
}

func (u *UnitOfWork) AttributeRepository() repository.AttributeRepository {
	return u.attributeRepositoryFactory.Create(u.tx)
}

func (u *UnitOfWork) ValueTypeRepository() repository.ValueTypeRepository {
	return u.valueTypeRepositoryFactory.Create(u.tx)
}
