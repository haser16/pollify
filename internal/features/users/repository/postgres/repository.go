package users_postgres_repository

import (
	core_postgres_pool "pollify/internal/core/repository/postgres/pool"
	core_redis "pollify/internal/core/repository/redis"
)

type UsersRepository struct {
	pool  core_postgres_pool.Pool
	redis core_redis.VerificationStore
}

func NewUsersRepository(
	pool core_postgres_pool.Pool,
	redis core_redis.VerificationStore,
) *UsersRepository {
	return &UsersRepository{
		pool:  pool,
		redis: redis,
	}
}
