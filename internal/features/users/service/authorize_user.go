package users_service

import (
	"context"
	"fmt"
	"pollify/internal/core/domain"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (u *UsersService) AuthorizeUser(
	ctx context.Context,
	data domain.UserAuthorize,
) (string, error) {
	userDomain, err := u.usersRepository.GetUserByEmail(ctx, data.Email)
	if err != nil {
		return "", fmt.Errorf("can't find user by email: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDomain.Password), []byte(data.Password))
	if err != nil {
		return "", fmt.Errorf("invalid password: %w", err)
	}

	token, err := generateJWTToken(userDomain.ID, u.jwtSecret)
	if err != nil {
		return "", fmt.Errorf("can't generate jwt token: %w", err)
	}

	return token, nil
}

func generateJWTToken(id int, secret []byte) (string, error) {
	claims := jwt.MapClaims{
		"id": id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
