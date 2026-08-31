package users_postgres_repository

import (
	"context"
	"fmt"
	"time"
)

func (r *UsersRepository) SaveVerificationToken(ctx context.Context, token string, userID int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	if err := r.redis.Set(ctx, token, userID, time.Minute*15); err != nil {
		return fmt.Errorf("save verification token: %w", err)
	}
	return nil
}
