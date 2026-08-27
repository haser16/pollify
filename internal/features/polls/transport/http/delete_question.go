package polls_transport_http

import (
	"net/http"
	core_logger "pollify/internal/core/logger"
	core_http_request "pollify/internal/core/transport/http/request"
	core_http_response "pollify/internal/core/transport/http/response"
)

func (h *PollsHTTPHandler) DeleteQuestion(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	questionID, err := core_http_request.GetIntPathValue(r, "question_id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get `question_id` from path",
		)
		return
	}
	if err := h.pollsService.DeleteQuestion(ctx, questionID); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to delete `question_id` from repository",
		)
		return
	}
	responseHandler.NoContentResponse()
}
