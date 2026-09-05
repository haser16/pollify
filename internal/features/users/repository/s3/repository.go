package users_s3_repository

import (
	core_s3 "pollify/internal/core/repository/s3"
)

type UsersRepository struct {
	storage core_s3.Storage
	bucket  string
}

func NewUsersRepository(
	storage core_s3.Storage,
	bucket string,
) *UsersRepository {
	return &UsersRepository{
		storage: storage,
		bucket:  bucket,
	}
}
