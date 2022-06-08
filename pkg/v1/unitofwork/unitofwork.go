package unitofwork

func NewUnitOfWork(tx Tx) UnitOfWork {
	return UnitOfWork{tx: tx}
}

type UnitOfWork struct {
	tx Tx
}

var _ Interface = (*UnitOfWork)(nil)

func (u *UnitOfWork) Complete(err *error) {
	var txErr error
	defer func() { *err = MergeErrors(err, &txErr) }()

	if *err == nil {
		txErr = u.tx.Commit() // TODO: liberr.Wrap(txErr, "cannot complete transaction")
	} else {
		txErr = u.tx.Rollback() // TODO: liberr.Wrap(txErr, "cannot rollback transaction"))
	}
}
