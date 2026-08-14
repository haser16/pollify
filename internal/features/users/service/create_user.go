package users_service

import (
	"context"
	"fmt"
	"pollify/internal/core/domain"

	"golang.org/x/crypto/bcrypt"
)

func (u *UsersService) CreateUser(
	ctx context.Context,
	user domain.User,
) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domain: %w", err)
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return domain.User{}, fmt.Errorf("hash user password: %w", err)
	}
	user.Password = hashedPassword

	userDomain, err := u.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user: %w", err)
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
