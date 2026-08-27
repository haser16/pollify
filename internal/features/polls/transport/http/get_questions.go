package polls_transport_http

import (
	"net/http"
	core_logger "pollify/internal/core/logger"
	core_http_request "pollify/internal/core/transport/http/request"
	core_http_response "pollify/internal/core/transport/http/response"
)

type GetQuestionsResponse []QuestionResponse

func (h *PollsHTTPHandler) GetQuestions(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	pollId, err := core_http_request.GetIntPathValue(r, "poll_id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get `id` from request",
		)
		return
	}

	questionsDomains, err := h.pollsService.GetQuestions(ctx, pollId)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get `questions`",
		)
	}
	response := GetQuestionsResponse(questionsDTOsFromDomains(questionsDomains))

	responseHandler.JsonResponse(response, http.StatusOK)
}
