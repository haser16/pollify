package votes_transport_http

import (
	"fmt"
	"net/http"
	core_logger "pollify/internal/core/logger"
	core_http_request "pollify/internal/core/transport/http/request"
	core_http_response "pollify/internal/core/transport/http/response"
)

type GetVotesResponse []VoteResponse

func (h *VotesHTTPHandler) GetVotes(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, questionID, optionID, err := getPathParam(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get path param",
		)
		return
	}
	votesDomains, err := h.votesService.GetVotes(ctx, userID, questionID, optionID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get votes",
		)
		return
	}
	response := GetVotesResponse(votesDTOsFromDomains(votesDomains))
	responseHandler.JsonResponse(response, http.StatusOK)
}

func getPathParam(r *http.Request) (*int, *int, *int, error) {
	userID, err := core_http_request.GetIntQueryParam(r, "user_id")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get userID query param: %w", err)
	}
	questionID, err := core_http_request.GetIntQueryParam(r, "question_id")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get questionID query param: %w", err)
	}
	optionID, err := core_http_request.GetIntQueryParam(r, "option_id")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to get optionID query param: %w", err)
	}
	return userID, questionID, optionID, nil
}
