package users_transport_http

import (
	"net/http"
	core_logger "pollify/internal/core/logger"
	core_http_request "pollify/internal/core/transport/http/request"
	core_http_response "pollify/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) EmailVerification(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	token, err := core_http_request.GetStringQueryParam(r, "token")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get token query param",
		)
		return
	}
	if err := h.usersService.EmailVerification(ctx, token); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to verify token",
		)
		return
	}
	responseHandler.JsonResponse("Success", http.StatusOK)
}
