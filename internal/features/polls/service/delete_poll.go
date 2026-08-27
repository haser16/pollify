package polls_service

import (
	"context"
	"fmt"
)

func (s *PollsService) DeletePoll(
	ctx context.Context,
	pollId int,
) error {
	if err := s.pollsRepository.DeletePoll(ctx, pollId); err != nil {
		return fmt.Errorf("failed to delete poll: %w", err)
	}
	return nil
}
