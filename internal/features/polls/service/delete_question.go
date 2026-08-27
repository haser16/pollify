package polls_service

import (
	"context"
	"fmt"
)

func (s *PollsService) DeleteQuestion(
	ctx context.Context,
	questionID int,
) error {
	if err := s.pollsRepository.DeleteQuestion(ctx, questionID); err != nil {
		return fmt.Errorf("can't delete question from repository: %w", err)
	}
	return nil
}
