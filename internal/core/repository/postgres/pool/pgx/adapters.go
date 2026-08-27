package core_pgx_pool

import (
	"context"
	"errors"
	"fmt"
	core_postgres_pool "pollify/internal/core/repository/postgres/pool"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgxRows struct {
	pgx.Rows
}

type pgxRow struct {
	pgx.Row
}

type pgxTx struct {
	pgx.Tx
}

func (r pgxRow) Scan(dest ...any) error {
	err := r.Row.Scan(dest...)
	if err != nil {
		return mapErrors(err)
	}
	return nil
}

func (t pgxTx) Commit(ctx context.Context) error {
	return t.Tx.Commit(ctx)
}

func (t pgxTx) Rollback(ctx context.Context) error {
	return t.Tx.Rollback(ctx)
}

type pgxCommandTag struct {
	pgconn.CommandTag
}

func mapErrors(err error) error {
	const (
		pgxViolatesForeignKeyErrorCode = "23503"
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return core_postgres_pool.ErrNoRows
	}
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == pgxViolatesForeignKeyErrorCode {
			return fmt.Errorf("%v: %w",
				err,
				core_postgres_pool.ErrViolatesForeignKey,
			)
		}
	}

	return fmt.Errorf(
		"%v: %w",
		err,
		core_postgres_pool.ErrUnknown,
	)
}
