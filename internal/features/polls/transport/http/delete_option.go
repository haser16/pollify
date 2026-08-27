package polls_transport_http

import (
	"net/http"
	core_logger "pollify/internal/core/logger"
	core_http_request "pollify/internal/core/transport/http/request"
	core_http_response "pollify/internal/core/transport/http/response"
)

func (h *PollsHTTPHandler) DeleteOption(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	optionID, err := core_http_request.GetIntPathValue(r, "option_id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get `option_id` path param",
		)
		return
	}
	if err := h.pollsService.DeleteOption(ctx, optionID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete option",
		)
		return
	}
	responseHandler.NoContentResponse()
}
