package polls_transport_http

import (
	"net/http"
	core_logger "pollify/internal/core/logger"
	core_http_request "pollify/internal/core/transport/http/request"
	core_http_response "pollify/internal/core/transport/http/response"
)

type PatchOptionRequest struct {
	OptionText string `json:"option_text"`
}

type PatchOptionResponse OptionResponse

func (h *PollsHTTPHandler) PatchOption(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request PatchOptionRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate request",
		)
		return
	}

	questionID, err := core_http_request.GetIntPathValue(r, "question_id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get question_id from path",
		)
		return
	}

	optionDomain := optionDomainFromDTO(request, questionID)
	optionDomain, err = h.pollsService.PatchOption(ctx, optionDomain)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch option",
		)
		return
	}
	response := PatchOptionResponse(optionDTOFromDomain(optionDomain))
	responseHandler.JsonResponse(response, http.StatusOK)
}
