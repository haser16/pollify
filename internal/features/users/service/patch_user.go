package users_service

import (
	"context"
	"fmt"
	"pollify/internal/core/domain"
)

func (u *UsersService) PatchUser(
	ctx context.Context,
	id int,
	patch domain.UserPatch,
) (domain.User, error) {
	user, err := u.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf(
			"failed to get user with id='%d': %w", id, err)
	}
	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("failed to apply patch: %w", err)
	}

	patchedUser, err := u.usersRepository.PatchUser(ctx, id, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to patch user: %w", err)
	}
	return patchedUser, nil
}
