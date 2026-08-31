package users_postgres_repository

import (
	"context"
	"fmt"
)

func (r *UsersRepository) VerifyEmail(
	ctx context.Context,
	userID int,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `UPDATE pollify.users
    SET email_verified = TRUE
    WHERE id = $1`

	result, err := r.pool.Exec(ctx, query, userID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("user %d not found", userID)
	}

	return nil
}
