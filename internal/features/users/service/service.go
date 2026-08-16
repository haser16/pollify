package users_service

import (
	"context"
	"pollify/internal/core/domain"
)

type UsersService struct {
	usersRepository UsersRepository
	jwtSecret       []byte
}

type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	GetUsers(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
	GetUser(
		ctx context.Context,
		id int,
	) (domain.User, error)
	DeleteUser(
		ctx context.Context,
		id int,
	) error
	PatchUser(
		ctx context.Context,
		id int,
		user domain.User,
	) (domain.User, error)
	GetUserByEmail(
		ctx context.Context,
		email string,
	) (domain.User, error)
}

func NewUsersService(usersRepository UsersRepository, jwtSecret string) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
		jwtSecret:       []byte(jwtSecret),
	}
}
