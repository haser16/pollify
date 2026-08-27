package votes_service

import (
	"context"
	core_domain "pollify/internal/core/domain"
)

type VotesService struct {
	votesRepository VotesRepository
}

type VotesRepository interface {
	CreateVote(
		ctx context.Context,
		domain core_domain.Vote,
	) (core_domain.Vote, error)
	GetVotes(
		ctx context.Context,
		userID *int,
		questionID *int,
		optionID *int,
	) ([]core_domain.Vote, error)
}

func NewVotesService(repository VotesRepository) *VotesService {
	return &VotesService{votesRepository: repository}
}
