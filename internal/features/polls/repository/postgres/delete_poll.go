package polls_postgres_repository

import (
	"context"
	"fmt"
	core_errors "pollify/internal/core/errors"
)

func (r *PollsRepository) DeletePoll(
	ctx context.Context,
	pollID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `DELETE FROM pollify.polls WHERE id = $1`

	cmdTag, err := r.pool.Exec(ctx, query, pollID)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("poll with id='%d' not found: %w", pollID, core_errors.ErrNotFound)
	}
	return nil
}
