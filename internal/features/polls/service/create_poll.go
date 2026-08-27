package polls_service

import (
	"context"
	"fmt"
	core_domain "pollify/internal/core/domain"
)

func (p *PollsService) CreatePoll(
	ctx context.Context,
	poll core_domain.Poll,
) (core_domain.Poll, error) {
	if err := poll.Validate(); err != nil {
		return core_domain.Poll{}, fmt.Errorf("failed to validate poll: %w", err)
	}
	pollDomain, err := p.pollsRepository.CreatePoll(ctx, poll)
	if err != nil {
		return core_domain.Poll{}, fmt.Errorf("failed to create poll: %w", err)
	}
	return pollDomain, nil
}
