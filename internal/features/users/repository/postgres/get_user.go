package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	"pollify/internal/core/domain"
	core_errors "pollify/internal/core/errors"
	core_postgres_pool "pollify/internal/core/repository/postgres/pool"
)

func (r *UsersRepository) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `SELECT id, full_name, email, phone_number FROM pollify.users WHERE id = $1;`
	row := r.pool.QueryRow(ctx, query, id)

	var userModel UserModel

	err := row.Scan(
		&userModel.ID,
		&userModel.FullName,
		&userModel.Email,
		&userModel.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"user with id='%d' is not exists: %w",
				id,
				core_errors.ErrNotFound,
			)
		}
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.FullName,
		userModel.Email,
		userModel.PhoneNumber,
		userModel.Password,
	)

	return userDomain, nil

}
