package polls_transport_http

import (
	"context"
	"net/http"
	core_domain "pollify/internal/core/domain"
	core_http_server "pollify/internal/core/transport/http/server"
)

type PollsHTTPHandler struct {
	pollsService PollsService
}

type PollsService interface {
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
		pollId int,
	) error
	GetQuestions(
		ctx context.Context,
		pollId int,
	) ([]core_domain.Question, error)
	DeleteQuestion(
		ctx context.Context,
		questionID int,
	) error
	GetOptions(
		ctx context.Context,
		questionID int,
	) ([]core_domain.Option, error)
	DeleteOption(
		ctx context.Context,
		optionID int,
	) error
	PatchOption(
		ctx context.Context,
		option core_domain.Option,
	) (core_domain.Option, error)
}

func NewPollsHTTPHandler(pollsService PollsService) *PollsHTTPHandler {
	return &PollsHTTPHandler{
		pollsService: pollsService,
	}
}

func (h *PollsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/polls",
			Handler: h.CreatePoll,
		},
		{
			Method:  http.MethodGet,
			Path:    "/polls",
			Handler: h.GetPolls,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/polls/{poll_id}",
			Handler: h.DeletePoll,
		},
		{
			Method:  http.MethodGet,
			Path:    "/polls/{poll_id}/questions",
			Handler: h.GetQuestions,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/polls/{poll_id}/questions/{question_id}",
			Handler: h.DeleteQuestion,
		},
		{
			Method:  http.MethodGet,
			Path:    "/polls/{poll_id}/questions/{question_id}/options/",
			Handler: h.GetOptions,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/polls/{poll_id}/questions/{question_id}/options/{option_id}",
			Handler: h.DeleteOption,
		},
	}
}
