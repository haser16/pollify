package polls_service

import (
	"context"
	"fmt"
	core_domain "pollify/internal/core/domain"
	core_errors "pollify/internal/core/errors"
)

func (s *PollsService) GetPolls(
	ctx context.Context,
	limit *int,
	offset *int,
) ([]core_domain.Poll, error) {
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

	polls, err := s.pollsRepository.GetPolls(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}
	return polls, nil
}
