package polls_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	core_domain "pollify/internal/core/domain"
	core_errors "pollify/internal/core/errors"

	"github.com/jackc/pgx/v5"
)

func (r *PollsRepository) GetOptions(
	ctx context.Context,
	pollId int,
) ([]core_domain.Option, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	checkQuestion := `SELECT id FROM pollify.questions WHERE id = $1`

	var id int
	err := r.pool.QueryRow(ctx, checkQuestion, pollId).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return []core_domain.Option{}, fmt.Errorf(
				"poll with id=`%d`: %w",
				pollId,
				core_errors.ErrNotFound,
			)
		}

		return []core_domain.Option{}, fmt.Errorf("scan error %w", err)
	}

	query := `SELECT id, question_id, option_text FROM pollify.options WHERE id = $1`
	rows, err := r.pool.Query(ctx, query, pollId)
	if err != nil {
		return []core_domain.Option{}, fmt.Errorf("get options: %w", pollId, err)
	}
	defer rows.Close()
	var options []core_domain.Option
	for rows.Next() {
		var option core_domain.Option
		if err := rows.Scan(&option.ID, &option.QuestionID, &option.OptionText); err != nil {
			return []core_domain.Option{}, fmt.Errorf("get options: %w", err)
		}
		options = append(options, option)
	}
	return options, nil
}
