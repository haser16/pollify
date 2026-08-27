package polls_transport_http

import (
	"fmt"
	"net/http"
	core_logger "pollify/internal/core/logger"
	core_http_request "pollify/internal/core/transport/http/request"
	core_http_response "pollify/internal/core/transport/http/response"
)

type GetPollsResponse []PollResponse

func (h *PollsHTTPHandler) GetPolls(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to parse http query params",
		)
		return
	}

	pollsDomain, err := h.pollsService.GetPolls(ctx, limit, offset)
	response := GetPollsResponse(pollsDTOsFromDomains(pollsDomain))

	responseHandler.JsonResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)

	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get 'limit' query param: %w", err)
	}
	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get 'offset' query param: %w", err)
	}
	return limit, offset, nil
}
