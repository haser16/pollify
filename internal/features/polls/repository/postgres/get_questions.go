package polls_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	core_domain "pollify/internal/core/domain"
	core_errors "pollify/internal/core/errors"

	"github.com/jackc/pgx/v5"
)

func (r *PollsRepository) GetQuestions(
	ctx context.Context,
	pollId int,
) ([]core_domain.Question, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	checkQuery := `
        SELECT id
        FROM pollify.polls
        WHERE id = $1;
    `

	var id int
	err := r.pool.QueryRow(ctx, checkQuery, pollId).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []core_domain.Question{}, fmt.Errorf(
				"poll with id=`%d`: %w",
				pollId,
				core_errors.ErrNotFound,
			)
		}

		return []core_domain.Question{}, fmt.Errorf("scan error %w", err)
	}

	query := `SELECT 
id,
poll_id,
question_text,
is_multiple
FROM pollify.questions
WHERE poll_id = $1
ORDER BY id ASC;
`
	rows, err := r.pool.Query(ctx, query, pollId)
	if err != nil {
		return []core_domain.Question{}, fmt.Errorf("scan error %w", err)
	}
	defer rows.Close()

	var questions []core_domain.Question

	for rows.Next() {
		var question core_domain.Question

		err := rows.Scan(
			&question.ID,
			&question.PollID,
			&question.QuestionText,
			&question.IsMultiple,
		)
		if err != nil {
			return []core_domain.Question{}, fmt.Errorf("scan error %w", err)
		}

		questions = append(questions, question)
	}

	return questions, nil
}
