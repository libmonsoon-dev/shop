package sqlunitofwork

import (
	"context"
	"database/sql"
)

type RepositoryFactory[Repo any] interface {
	Create(db Tx) Repo
}

type Tx interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}
