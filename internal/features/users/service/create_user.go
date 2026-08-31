package users_service

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	core_domain "pollify/internal/core/domain"
	core_publisher "pollify/internal/core/publisher"

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

	verificationToken, err := generateVerificationToken()
	if err != nil {
		return core_domain.User{}, fmt.Errorf("generate verification token: %w", err)
	}

	if err := u.usersRepository.SaveVerificationToken(ctx, verificationToken, userDomain.ID); err != nil {
		return core_domain.User{}, fmt.Errorf("save verification token: %w", err)
	}

	if err := u.publisher.Publish(ctx,
		core_publisher.VerificationMessage{
			Email: user.Email,
			Token: verificationToken,
		}); err != nil {
		return core_domain.User{}, fmt.Errorf("publish user: %w", err)
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

func generateVerificationToken() (string, error) {
	b := make([]byte, 32)

	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
