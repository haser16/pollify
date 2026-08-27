package polls_postgres_repository

import (
	"context"
	"fmt"
	core_errors "pollify/internal/core/errors"
)

func (r *PollsRepository) DeleteQuestion(
	ctx context.Context,
	questionId int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := "DELETE FROM pollify.question WHERE id = $1"
	cmdTag, err := r.pool.Exec(ctx, query, questionId)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("question with id='%d' not found: %w", questionId, core_errors.ErrNotFound)
	}
	return nil
}
