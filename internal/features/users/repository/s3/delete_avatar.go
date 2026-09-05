package users_s3_repository

import "context"

func (r *UsersRepository) DeleteAvatar(
	ctx context.Context,
	key string,
) error {
	return r.storage.Delete(
		ctx,
		r.bucket,
		key,
	)
}
