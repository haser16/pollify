package polls_transport_http

import (
	"net/http"
	core_logger "pollify/internal/core/logger"
	core_http_response "pollify/internal/core/transport/http/response"
)

func (h *PollsHTTPHandler) GetPolls(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

}

func getLimitOffsetQueryParams(r *http.Request) (limit, offset int, err error) {}
