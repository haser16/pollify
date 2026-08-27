package users_service

import (
	"context"
	"fmt"
	core_domain "pollify/internal/core/domain"

	"golang.org/x/crypto/bcrypt"
)

func (u *UsersService) CreateUser(
	ctx context.Context,
	user core_domain.User,
) (core_domain.User, error) {
	if err := user.Validate(); err != nil {
		return core_domain.User{}, fmt.Errorf("validate user core_domain: %w", err)
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return core_domain.User{}, fmt.Errorf("hash user password: %w", err)
	}
	user.Password = hashedPassword

	userDomain, err := u.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return core_domain.User{}, fmt.Errorf("create user: %w", err)
	}
	return userDomain, nil
}

func hashPassword(password string) (string, error) {
	passwordBytes := []byte(password)

	hashedBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hash password: %w", err)
	}
	return string(hashedBytes), nil
}
