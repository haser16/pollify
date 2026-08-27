package polls_transport_http

import (
	"net/http"
	core_logger "pollify/internal/core/logger"
	core_http_request "pollify/internal/core/transport/http/request"
	core_http_response "pollify/internal/core/transport/http/response"
)

type CreatePollRequest PollRequest

func (h *PollsHTTPHandler) CreatePoll(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreatePollRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate request",
		)
		return
	}

	pollDomain := domainFromDTO(request)

	var err error
	pollDomain, err = h.pollsService.CreatePoll(ctx, pollDomain)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to create poll",
		)
		return
	}

	response := pollDTOFromDomain(pollDomain)
	responseHandler.JsonResponse(response, http.StatusCreated)
}
