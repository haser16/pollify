package polls_service

import (
	"context"
	"fmt"
	core_domain "pollify/internal/core/domain"
)

func (s *PollsService) GetQuestions(
	ctx context.Context,
	pollId int,
) ([]core_domain.Question, error) {
	questions, err := s.pollsRepository.GetQuestions(ctx, pollId)
	if err != nil {
		return nil, fmt.Errorf("failed to get qeustions by id: %w", err)
	}
	return questions, nil
}
