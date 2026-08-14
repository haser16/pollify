package users_transport_http

import (
	"net/http"
	"pollify/internal/core/domain"
	core_logger "pollify/internal/core/logger"
	core_http_request "pollify/internal/core/transport/http/request"
	core_http_response "pollify/internal/core/transport/http/response"
	core_http_types "pollify/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	FullName core_http_types.Nullable[string] `json:"full_name"`
	Email    core_http_types.Nullable[string] `json:"email"`
	Phone    core_http_types.Nullable[string] `json:"phone_number"`
}

type PatchUserResponse UserDTOResponse

func (h *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"get `UserID` path value",
		)
		return
	}

	var patchRequest PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &patchRequest); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to decode and validate patch request",
		)
		return
	}

	userPatch := userPatchFromRequest(patchRequest)
	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to patch user",
		)
	}
	response := PatchUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JsonResponse(response, http.StatusOK)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		request.FullName.ToDomain(),
		request.Email.ToDomain(),
		request.Phone.ToDomain(),
	)
}
