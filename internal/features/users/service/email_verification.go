package users_service

import (
	"context"
	"fmt"
)

func (s *UsersService) EmailVerification(
	ctx context.Context,
	token string,
) error {
	userID, err := s.usersRepository.GetUserByToken(ctx, token)
	if err != nil {
		return fmt.Errorf("failed to get user token from redis: %w", err)
	}
	if err := s.usersRepository.VerifyEmail(ctx, userID); err != nil {
		return fmt.Errorf("failed to verify email: %w", err)
	}
	return nil
}
