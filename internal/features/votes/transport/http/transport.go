package votes_transport_http

import (
	"context"
	"net/http"
	core_domain "pollify/internal/core/domain"
	core_http_server "pollify/internal/core/transport/http/server"
)

type VotesHTTPHandler struct {
	votesService VotesService
}

type VotesService interface {
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

func NewVotesHTTPHandler(service VotesService) *VotesHTTPHandler {
	return &VotesHTTPHandler{votesService: service}
}

func (h *VotesHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodGet,
			Path:    "/votes",
			Handler: h.GetVotes,
		},
		{
			Method:  http.MethodPost,
			Path:    "/votes",
			Handler: h.CreateVote,
		},
	}
}
