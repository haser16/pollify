package users_s3_repository

import (
	"context"
	"fmt"
	"io"
)

func (r *UsersRepository) UploadAvatar(
	ctx context.Context,
	userID string,
	file io.Reader,
	contentType string,
) error {

	key := fmt.Sprintf(
		"users/%s/avatar",
		userID,
	)

	return r.storage.Upload(
		ctx,
		r.bucket,
		key,
		file,
		contentType,
	)
}
