package polls_postgres_repository

import (
	"context"
	"fmt"
	core_errors "pollify/internal/core/errors"
)

func (r *PollsRepository) DeleteOption(
	ctx context.Context,
	optionId int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `DELETE FROM pollify.options WHERE id = $1`

	cmdTag, err := r.pool.Exec(ctx, query, optionId)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("option with id='%d' not found: %w", optionId, core_errors.ErrNotFound)
	}
	return nil
}
