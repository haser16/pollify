package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (r *UsersRepository) GetUserByToken(
	ctx context.Context,
	token string,
) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()
	
	userID, err := r.redis.Get(ctx, token)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, fmt.Errorf("user not found")
		}

		return 0, err
	}

	return int(userID), nil
}
