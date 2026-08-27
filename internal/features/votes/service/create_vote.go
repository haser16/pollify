package votes_service

import (
	"context"
	"fmt"
	core_domain "pollify/internal/core/domain"
)

func (s *VotesService) CreateVote(
	ctx context.Context,
	vote core_domain.Vote,
) (core_domain.Vote, error) {
	vote, err := s.votesRepository.CreateVote(ctx, vote)
	if err != nil {
		return core_domain.Vote{}, fmt.Errorf("repository failed: %w", err)
	}
	return vote, nil
}
