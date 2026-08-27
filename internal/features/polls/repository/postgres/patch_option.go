package polls_postgres_repository

import (
	"context"
	"errors"
	"fmt"
	core_domain "pollify/internal/core/domain"
	core_errors "pollify/internal/core/errors"
	core_postgres_pool "pollify/internal/core/repository/postgres/pool"
)

func (r *PollsRepository) PatchOption(
	ctx context.Context,
	option core_domain.Option,
) (core_domain.Option, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTimeOut())
	defer cancel()

	query := "UPDATE pollify.options SET option_text=$1 WHERE id=$2 RETURNING id, question_id, option_text;"

	row := r.pool.QueryRow(ctx, query, option.OptionText, option.ID)
	var optionModel core_domain.Option

	err := row.Scan(&optionModel.ID, &optionModel.QuestionID, &optionModel.OptionText)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return core_domain.Option{}, fmt.Errorf(
				"option with id='%d' concurrently accessed: %w",
				option.ID,
				core_errors.ErrConflict,
			)
		}
		return core_domain.Option{}, fmt.Errorf("scan error: %w", err)
	}

	optionDomain := core_domain.NewOption(
		optionModel.ID,
		optionModel.QuestionID,
		optionModel.OptionText,
	)
	return optionDomain, nil

}
