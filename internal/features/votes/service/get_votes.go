package votes_service

import (
	"context"
	"fmt"
	core_domain "pollify/internal/core/domain"
)

func (s *VotesService) GetVotes(
	ctx context.Context,
	userID *int,
	questionID *int,
	optionID *int,
) ([]core_domain.Vote, error) {
	polls, err := s.votesRepository.GetVotes(ctx, userID, questionID, optionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get votes: %w", err)
	}
	return polls, nil
}
