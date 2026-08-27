package votes_postgres_repository

import (
	"context"
	core_domain "pollify/internal/core/domain"
)

func (r *VotesRepository) GetVotes(
	ctx context.Context,
	userID *int,
	questionID *int,
	optionID *int,
) ([]core_domain.Vote, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `
		SELECT
			id,
			user_id,
			question_id,
			option_id,
			voted_at
		FROM pollify.votes
		WHERE ($1::int IS NULL OR user_id = $1)
		  AND ($2::int IS NULL OR question_id = $2)
		  AND ($3::int IS NULL OR option_id = $3)
		ORDER BY id
	`

	rows, err := r.pool.Query(
		ctx,
		query,
		userID,
		questionID,
		optionID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	votes := make([]core_domain.Vote, 0)

	for rows.Next() {
		var vote core_domain.Vote

		err := rows.Scan(
			&vote.ID,
			&vote.UserID,
			&vote.QuestionID,
			&vote.OptionID,
			&vote.VotedAt,
		)
		if err != nil {
			return nil, err
		}

		votes = append(votes, vote)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return votes, nil
}
