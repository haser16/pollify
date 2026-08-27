package polls_service

import (
	"context"
	"fmt"
)

func (s *PollsService) DeleteOption(
	ctx context.Context,
	optionID int,
) error {
	if err := s.pollsRepository.DeleteOption(ctx, optionID); err != nil {
		return fmt.Errorf("can't delete option with id='%d': %w", optionID, err)
	}
	return nil
}
