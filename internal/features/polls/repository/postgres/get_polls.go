package polls_postgres_repository

import (
	"context"
	"fmt"
	core_domain "pollify/internal/core/domain"
)

func (r *PollsRepository) GetPolls(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]core_domain.Poll, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `SELECT 
id, title, description, created_at, expires_at, completed, owner_id
FROM pollify.polls
ORDER BY id ASC LIMIT $1 OFFSET $2;`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get polls: %w", err)
	}
	defer rows.Close()

	var polls []core_domain.Poll
	for rows.Next() {
		var poll core_domain.Poll
		err = rows.Scan(
			&poll.ID,
			&poll.Title,
			&poll.Description,
			&poll.CreatedAt,
			&poll.ExpiresAt,
			&poll.Completed,
			&poll.AuthorID,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan poll: %w", err)
		}
		polls = append(polls, poll)
	}
	return polls, nil
}
