package polls_service

import (
	"context"
	"fmt"
	core_domain "pollify/internal/core/domain"
)

func (s *PollsService) GetOptions(
	ctx context.Context,
	questionID int,
) ([]core_domain.Option, error) {
	options, err := s.pollsRepository.GetOptions(ctx, questionID)
	if err != nil {
		return nil, fmt.Errorf("get options: %w", err)
	}
	return options, nil
}
