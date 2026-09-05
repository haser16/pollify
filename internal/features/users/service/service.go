package users_service

import (
	"context"
	"io"
	"pollify/internal/core/domain"
	"pollify/internal/core/publisher"
)

type UsersService struct {
	usersRepository UsersRepository
	publisher       core_publisher.Publisher
	s3              UsersS3Repository

	jwtSecret []byte
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
	SaveVerificationToken(
		ctx context.Context,
		token string,
		userID int,
	) error
	GetUserByToken(
		ctx context.Context,
		token string,
	) (int, error)
	VerifyEmail(
		ctx context.Context,
		userID int,
	) error
}

type UsersS3Repository interface {
	UploadAvatar(
		ctx context.Context,
		userID string,
		file io.Reader,
		contentType string,
	) error
}

func NewUsersService(
	usersRepository UsersRepository,
	jwtSecret string,
	publisher core_publisher.Publisher,
	s3 UsersS3Repository,
) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
		jwtSecret:       []byte(jwtSecret),
		publisher:       publisher,
		s3:              s3,
	}
}
