package polls_transport_http

import (
	core_domain "pollify/internal/core/domain"
	"time"
)

type PollRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ExpiresAt   int64  `json:"expires_at"`
	AuthorID    int    `json:"author_id"`
	Questions   []QuestionRequest
}

type QuestionRequest struct {
	Text       string `json:"question_text"`
	IsMultiple bool   `json:"is_multiple"`
	Options    []OptionRequest
}
type OptionRequest struct {
	Text string `json:"option_text"`
}

type PollResponse struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	Completed   bool      `json:"completed"`
	AuthorID    int       `json:"author_id"`
}

type QuestionResponse struct {
	ID           int    `json:"id"`
	PollID       int    `json:"poll_id"`
	QuestionText string `json:"question_text"`
	IsMultiple   bool   `json:"is_multiple"`
}

type OptionResponse struct {
	ID         int    `json:"id"`
	QuestionID int    `json:"question_id"`
	OptionText string `json:"option_text"`
}

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

func domainFromDTO(request CreatePollRequest) core_domain.Poll {
	domainQuestions := make([]core_domain.Question, len(request.Questions))

	for i, q := range request.Questions {
		domainOptions := make([]core_domain.Option, len(q.Options))
		for j, o := range q.Options {
			domainOptions[j] = core_domain.Option{
				ID:         core_domain.UnInitializedID,
				QuestionID: core_domain.UnInitializedID,
				OptionText: o.Text,
			}
		}

		domainQuestions[i] = core_domain.Question{
			ID:           core_domain.UnInitializedID,
			PollID:       core_domain.UnInitializedID,
			QuestionText: q.Text,
			IsMultiple:   q.IsMultiple,
			Options:      domainOptions,
		}
	}

	return core_domain.NewPollUninitialized(
		request.Title,
		request.Description,
		time.UnixMilli(request.ExpiresAt),
		request.AuthorID,
		domainQuestions,
	)
}

func pollDTOFromDomain(poll core_domain.Poll) PollResponse {
	return PollResponse{
		ID:          poll.ID,
		Title:       poll.Title,
		Description: poll.Description,
		CreatedAt:   poll.CreatedAt,
		ExpiresAt:   poll.ExpiresAt,
		Completed:   poll.Completed,
		AuthorID:    poll.AuthorID,
	}
}

func pollsDTOsFromDomains(polls []core_domain.Poll) []PollResponse {
	responses := make([]PollResponse, len(polls))

	for i, p := range polls {
		responses[i] = pollDTOFromDomain(p)
	}

	return responses
}

func questionDTOFromDomain(question core_domain.Question) QuestionResponse {
	return QuestionResponse{
		ID:           question.ID,
		PollID:       question.PollID,
		QuestionText: question.QuestionText,
		IsMultiple:   question.IsMultiple,
	}
}

func questionsDTOsFromDomains(questions []core_domain.Question) []QuestionResponse {
	responses := make([]QuestionResponse, len(questions))
	for i, p := range questions {
		responses[i] = questionDTOFromDomain(p)
	}
	return responses
}

func optionDTOFromDomain(option core_domain.Option) OptionResponse {
	return OptionResponse{
		ID:         option.ID,
		QuestionID: option.QuestionID,
		OptionText: option.OptionText,
	}
}

func optionsDTOsFromDomain(options []core_domain.Option) []OptionResponse {
	responses := make([]OptionResponse, len(options))
	for i, p := range options {
		responses[i] = optionDTOFromDomain(p)
	}
	return responses
}

func optionDomainFromDTO(option PatchOptionRequest, questionID int) core_domain.Option {
	return core_domain.NewOptionUninitialized(questionID, option.OptionText)
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

func voteDomainFromDTO(vote CreateVoteRequest) core_domain.Vote {
	return core_domain.NewVoteUninitialized(vote.UserID, vote.QuestionID, vote.OptionID)
}
