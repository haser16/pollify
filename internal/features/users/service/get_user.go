package users_service

import (
	"context"
	"fmt"
	"pollify/internal/core/domain"
)

func (u *UsersService) GetUser(
	ctx context.Context,
	id int,
) (domain.User, error) {
	user, err := u.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}
	return user, nil
}
