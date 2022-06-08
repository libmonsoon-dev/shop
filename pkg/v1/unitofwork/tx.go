package unitofwork

type Tx interface {
	Commit() error
	Rollback() error
}
