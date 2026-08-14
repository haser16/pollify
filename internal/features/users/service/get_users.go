package users_service

import (
	"context"
	"fmt"
	"pollify/internal/core/domain"
	core_errors "pollify/internal/core/errors"
)

func (u *UsersService) GetUsers(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil,
			fmt.Errorf(
				"limit must be a non-negative integer: %w",
				core_errors.ErrInvalidArgument,
			)
	}

	if offset != nil && *offset < 0 {
		return nil,
			fmt.Errorf(
				"offset must be a non-negative integer: %w",
				core_errors.ErrInvalidArgument,
			)
	}

	users, err := u.usersRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}
	return users, nil
}
