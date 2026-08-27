package polls_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	core_domain "pollify/internal/core/domain"
	core_errors "pollify/internal/core/errors"
	core_postgres_pool "pollify/internal/core/repository/postgres/pool"
)

func (r *PollsRepository) CreatePoll(
	ctx context.Context,
	poll core_domain.Poll,
) (core_domain.Poll, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return core_domain.Poll{}, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	pollQuery := `
		INSERT INTO pollify.polls (title, description, created_at, expires_at, completed, owner_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at;`

	var insertedPollID int

	err = tx.QueryRow(
		ctx,
		pollQuery,
		poll.Title,
		poll.Description,
		poll.CreatedAt,
		poll.ExpiresAt,
		poll.Completed,
		poll.AuthorID,
	).Scan(&insertedPollID, &poll.CreatedAt)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
			return core_domain.Poll{}, fmt.Errorf(
				"%v: user with id=`%d`: %w",
				err,
				poll.AuthorID,
				core_errors.ErrNotFound,
			)
		}

		return core_domain.Poll{}, fmt.Errorf("scan error %w", err)
	}
	poll.ID = insertedPollID

	savedQuestions := make([]core_domain.Question, len(poll.Questions))

	for i, q := range poll.Questions {
		questionQuery := `
			INSERT INTO pollify.questions (poll_id, question_text, is_multiple)
			VALUES ($1, $2, $3)
			RETURNING id;`

		var insertedQuestionID int
		err = tx.QueryRow(ctx, questionQuery, poll.ID, q.QuestionText, q.IsMultiple).Scan(&insertedQuestionID)
		if err != nil {
			return core_domain.Poll{}, fmt.Errorf("failed to insert question at index %d: %w", i, err)
		}

		var savedOptions []core_domain.Option
		optionQuery := `
			INSERT INTO pollify.options (question_id, option_text)
			VALUES ($1, $2)
			RETURNING id, question_id, option_text;`

		for _, opt := range q.Options {
			var insertedOpt core_domain.Option

			err = tx.QueryRow(ctx, optionQuery, insertedQuestionID, opt.OptionText).
				Scan(&insertedOpt.ID, &insertedOpt.QuestionID, &insertedOpt.OptionText)

			if err != nil {
				return core_domain.Poll{}, fmt.Errorf("failed to insert option '%s': %w", opt.OptionText, err)
			}
			savedOptions = append(savedOptions, insertedOpt)
		}

		savedQuestions[i] = core_domain.Question{
			ID:           insertedQuestionID,
			PollID:       poll.ID,
			QuestionText: q.QuestionText,
			IsMultiple:   q.IsMultiple,
			Options:      savedOptions,
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return core_domain.Poll{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	poll.Questions = savedQuestions
	return poll, nil
}
