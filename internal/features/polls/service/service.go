package polls_service

import (
	"context"
	core_domain "pollify/internal/core/domain"
)

type PollsService struct {
	pollsRepository PollsRepository
}

type PollsRepository interface {
	CreatePoll(
		ctx context.Context,
		poll core_domain.Poll,
	) (core_domain.Poll, error)
	GetPolls(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]core_domain.Poll, error)
	DeletePoll(
		ctx context.Context,
		pollID int,
	) error
	GetQuestions(
		ctx context.Context,
		pollId int,
	) ([]core_domain.Question, error)
	DeleteQuestion(
		ctx context.Context,
		questionId int,
	) error
	GetOptions(
		ctx context.Context,
		pollId int,
	) ([]core_domain.Option, error)
	DeleteOption(
		ctx context.Context,
		optionID int,
	) error
	PatchOption(
		ctx context.Context,
		option core_domain.Option,
	) (core_domain.Option, error)
	GetVotes(
		ctx context.Context,
		userID *int,
		questionID *int,
		optionID *int,
	) ([]core_domain.Vote, error)
	CreateVote(
		ctx context.Context,
		vote core_domain.Vote,
	) (core_domain.Vote, error)
}

func NewPollsService(poolsRepository PollsRepository) *PollsService {
	return &PollsService{
		pollsRepository: poolsRepository,
	}
}
