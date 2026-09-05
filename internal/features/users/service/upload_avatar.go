package users_service

import (
	"context"
	"fmt"
	"io"
)

func (u *UsersService) UploadAvatar(
	ctx context.Context,
	userID string,
	file io.Reader,
	contentType string,
) error {
	if err := u.s3.UploadAvatar(
		ctx,
		userID,
		file,
		contentType,
	); err != nil {
		return fmt.Errorf("upload user avatar: %w", err)
	}

	return nil
}
