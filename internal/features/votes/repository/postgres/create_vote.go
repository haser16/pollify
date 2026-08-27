package votes_postgres_repository

import (
	"context"
	"fmt"
	core_domain "pollify/internal/core/domain"
)

func (r *VotesRepository) CreateVote(
	ctx context.Context,
	vote core_domain.Vote,
) (core_domain.Vote, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `INSERT INTO pollify.votes (user_id, question_id, option_id) VALUES ($1, $2, $3)
RETURNING id, user_id, question_id, option_id`
	row := r.pool.QueryRow(ctx, query, vote.UserID, vote.QuestionID, vote.OptionID)

	var voteResponse core_domain.Vote
	err := row.Scan(&voteResponse.ID, &voteResponse.UserID, &voteResponse.QuestionID, &voteResponse.OptionID)
	if err != nil {
		return core_domain.Vote{}, fmt.Errorf("error creating vote: %w", err)
	}
	return voteResponse, nil
}
