package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"pollify/internal/core/domain"
	core_errors "pollify/internal/core/errors"
	core_postgres_pool "pollify/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) PatchUser(
	ctx context.Context,
	id int,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `
	UPDATE pollify.users
	SET 
		full_name=$1,
		email=$2,
		phone_number=$3
	WHERE id=$4
	RETURNING
		 id,
		 full_name,
		 email,
		 phone_number;
	`

	row := r.pool.QueryRow(
		ctx,
		query,
		user.FullName,
		user.Email,
		user.PhoneNumber,
		user.ID,
	)
	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.Email,
		&userModel.FullName,
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id='%d' concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Email,
		userModel.FullName,
		userModel.PhoneNumber,
	)
	return userDomain, nil
}
