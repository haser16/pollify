package users_postgres_repository

import (
	"context"
	"fmt"
	"pollify/internal/core/domain"
)

func (r *UsersRepository) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := `
		INSERT INTO pollify.users
		(full_name, email, phone_number, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id, full_name, email, phone_number`

	row := r.pool.QueryRow(ctx, query, user.FullName, user.Email, user.PhoneNumber, user.Password)

	var userModel UserModel
	err := row.Scan(
		&userModel.ID,
		&userModel.FullName,
		&userModel.Email,
		&userModel.PhoneNumber,
	)
	if err != nil {
		return domain.User{}, fmt.Errorf("scan error: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.FullName,
		userModel.Email,
		userModel.PhoneNumber,
	)

	return userDomain, nil
}
