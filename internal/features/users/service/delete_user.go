package users_service

import (
	"context"
	"fmt"
)

func (u *UsersService) DeleteUser(
	ctx context.Context,
	id int,
) error {
	if err := u.usersRepository.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("delete user from repository: %w", err)
	}
	return nil
}
