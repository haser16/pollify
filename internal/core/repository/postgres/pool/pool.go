package core_postgres_pool

import (
	"context"
	"time"
)

type Pool interface {
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row
	Exec(ctx context.Context, sql string, arguments ...any) (CommandTag, error)

	Close()
	OperationTimeOut() time.Duration

	Begin(ctx context.Context) (Tx, error)
}

type Rows interface {
	Close()
	Err() error
	Next() bool
	Scan(dest ...any) error
}

type Row interface {
	Scan(dest ...any) error
}

type CommandTag interface {
	RowsAffected() int64
}

type Tx interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error

	Exec(ctx context.Context, sql string, arguments ...any) (commandTag CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row
}
