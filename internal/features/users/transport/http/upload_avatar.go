package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "pollify/internal/core/logger"
	core_http_response "pollify/internal/core/transport/http/response"
)

const maxAvatarSize = 5 << 20 // 5 MB

func (h *UsersHTTPHandler) UploadAvatar(
	rw http.ResponseWriter,
	r *http.Request,
) {
	ctx := r.Context()

	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID := r.PathValue("id")
	if userID == "" {
		responseHandler.ErrorResponse(
			fmt.Errorf("user id is required"),
			"failed to upload avatar",
		)
		return
	}

	r.Body = http.MaxBytesReader(
		rw,
		r.Body,
		maxAvatarSize,
	)

	if err := r.ParseMultipartForm(maxAvatarSize); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to parse multipart form",
		)
		return
	}

	file, header, err := r.FormFile("avatar")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"avatar is required",
		)
		return
	}
	defer file.Close()

	contentType := header.Header.Get("Content-Type")

	if err := h.usersService.UploadAvatar(
		ctx,
		userID,
		file,
		contentType,
	); err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to upload avatar",
		)
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
