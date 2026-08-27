package votes_transport_http

import (
	core_domain "pollify/internal/core/domain"
	"time"
)

type VoteResponse struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	QuestionID int       `json:"question_id"`
	OptionID   int       `json:"option_id"`
	VotedAt    time.Time `json:"voted_at"`
}

type VoteRequest struct {
	UserID     int `json:"user_id"`
	QuestionID int `json:"question_id"`
	OptionID   int `json:"option_id"`
}

func voteDTOFromDomain(vote core_domain.Vote) VoteResponse {
	return VoteResponse{
		ID:         vote.ID,
		UserID:     vote.UserID,
		QuestionID: vote.QuestionID,
		OptionID:   vote.OptionID,
		VotedAt:    vote.VotedAt,
	}
}

func votesDTOsFromDomains(votes []core_domain.Vote) []VoteResponse {
	responses := make([]VoteResponse, len(votes))
	for i, p := range votes {
		responses[i] = voteDTOFromDomain(p)
	}
	return responses
}

func voteDomainFromDTO(request CreateVoteRequest) core_domain.Vote {
	return core_domain.NewVoteUninitialized(request.UserID, request.QuestionID, request.OptionID)
}
