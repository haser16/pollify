package users_transport_http

import (
	"net/http"
	"pollify/internal/core/domain"
	core_logger "pollify/internal/core/logger"
	core_http_request "pollify/internal/core/transport/http/request"
	core_http_response "pollify/internal/core/transport/http/response"
)

type AuthorizeUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthorizeUserResponse struct {
	JWT string `json:"jwt"`
}

func (h *UsersHTTPHandler) AuthorizeUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request AuthorizeUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate AuthorizeUser request",
		)
		return
	}
	authorizeDomain := domainFromDTOAuthorize(request)

	jwt, err := h.usersService.AuthorizeUser(ctx, authorizeDomain)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to authorize user",
		)
		return
	}
	response := AuthorizeUserResponse{JWT: jwt}
	responseHandler.JsonResponse(response, http.StatusOK)
}

func domainFromDTOAuthorize(request AuthorizeUserRequest) domain.UserAuthorize {
	return domain.NewAuthorizeUser(request.Email, request.Password)
}
