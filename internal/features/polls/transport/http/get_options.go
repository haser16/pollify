package polls_transport_http

import (
	"net/http"
	core_logger "pollify/internal/core/logger"
	core_http_request "pollify/internal/core/transport/http/request"
	core_http_response "pollify/internal/core/transport/http/response"
)

type GetOptionsResponse []OptionResponse

func (h *PollsHTTPHandler) GetOptions(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	questionID, err := core_http_request.GetIntPathValue(r, "question_id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get `question_id` from request",
		)
		return
	}

	optionsDomains, err := h.pollsService.GetOptions(ctx, questionID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get `options` from database",
		)
		return
	}
	response := GetOptionsResponse(optionsDTOsFromDomain(optionsDomains))
	responseHandler.JsonResponse(response, http.StatusOK)
}
