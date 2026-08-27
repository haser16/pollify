package polls_transport_http

import (
	"net/http"
	core_logger "pollify/internal/core/logger"
	core_http_request "pollify/internal/core/transport/http/request"
	core_http_response "pollify/internal/core/transport/http/response"
)

func (h *PollsHTTPHandler) DeletePoll(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	pollID, err := core_http_request.GetIntPathValue(r, "poll_id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get `poll_id` path parameter",
		)
		return
	}
	if err := h.pollsService.DeletePoll(ctx, pollID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete poll from repository",
		)
		return
	}
	responseHandler.NoContentResponse()
}
