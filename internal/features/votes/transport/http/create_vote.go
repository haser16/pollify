package votes_transport_http

import (
	"net/http"
	core_logger "pollify/internal/core/logger"
	core_http_request "pollify/internal/core/transport/http/request"
	core_http_response "pollify/internal/core/transport/http/response"
)

type CreateVoteRequest VoteRequest
type CreateVoteResponse VoteResponse

func (h *VotesHTTPHandler) CreateVote(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateVoteRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate request",
		)
		return
	}

	voteDomain := voteDomainFromDTO(request)

	var err error
	voteDomain, err = h.votesService.CreateVote(ctx, voteDomain)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create vote",
		)
		return
	}

	response := CreateVoteResponse(voteDTOFromDomain(voteDomain))
	responseHandler.JsonResponse(response, http.StatusCreated)

}
