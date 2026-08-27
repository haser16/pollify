package polls_service

import (
	"context"
	"fmt"
	core_domain "pollify/internal/core/domain"
)

func (s *PollsService) PatchOption(
	ctx context.Context,
	option core_domain.Option,
) (core_domain.Option, error) {
	option, err := s.pollsRepository.PatchOption(ctx, option)
	if err != nil {
		return option, fmt.Errorf("failed to patch option repository: %w", err)
	}
	return option, nil
}
