package domain

import (
	"fmt"
	core_errors "pollify/internal/core/errors"
	"time"
)

type Poll struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	ExpiresAt   time.Time
	Completed   bool
	AuthorID    int
	Questions   []Question
}

type Question struct {
	ID           int
	PollID       int
	QuestionText string
	IsMultiple   bool
	Options      []Option
}

type Option struct {
	ID         int
	QuestionID int
	OptionText string
}

type Vote struct {
	ID         int
	UserID     int
	QuestionID int
	OptionID   int
	VotedAt    time.Time
}

func NewPoll(
	id int,
	title string,
	description string,
	createdAt time.Time,
	expiresAt time.Time,
	completed bool,
	authorID int,
	questions []Question,
) Poll {
	return Poll{
		ID:          id,
		Title:       title,
		Description: description,
		CreatedAt:   createdAt,
		ExpiresAt:   expiresAt,
		Completed:   completed,
		AuthorID:    authorID,
		Questions:   questions,
	}
}

func NewPollUninitialized(
	title string,
	description string,
	expiresAt time.Time,
	authorID int,
	questions []Question,
) Poll {
	return NewPoll(
		UnInitializedID,
		title,
		description,
		time.Now(),
		expiresAt,
		false,
		authorID,
		questions,
	)
}

func NewVote(
	id int,
	userID int,
	questionID int,
	optionID int,
	votedAt time.Time,
) Vote {
	return Vote{
		ID:         id,
		UserID:     userID,
		QuestionID: questionID,
		OptionID:   optionID,
		VotedAt:    votedAt,
	}
}

func NewVoteUninitialized(
	userID int,
	questionID int,
	optionID int,
) Vote {
	return NewVote(
		UnInitializedID,
		userID,
		questionID,
		optionID,
		time.Now(),
	)
}

func NewOption(
	id int,
	questionID int,
	optionText string,
) Option {
	return Option{
		ID:         id,
		QuestionID: questionID,
		OptionText: optionText,
	}
}

func NewOptionUninitialized(questionID int, optionText string) Option {
	return Option{
		ID:         UnInitializedID,
		QuestionID: questionID,
		OptionText: optionText,
	}
}

func (p *Poll) Validate() error {
	titleLen := len([]rune(p.Title))
	if titleLen < 2 || titleLen > 100 {
		return fmt.Errorf("invalid poll `Title` length: %d: %w", titleLen, core_errors.ErrInvalidArgument)
	}
	descriptionLen := len([]rune(p.Description))
	if descriptionLen < 2 || descriptionLen > 1000 {
		return fmt.Errorf("invalid poll `Description` length: %d: %w", descriptionLen, core_errors.ErrInvalidArgument)
	}
	if p.ExpiresAt.Before(p.CreatedAt) {
		return fmt.Errorf("invalid poll `ExpiresAt: %v: %w", p.ExpiresAt, core_errors.ErrInvalidArgument)
	}

	if len(p.Questions) == 0 {
		return fmt.Errorf("poll must have at least one question: %w", core_errors.ErrInvalidArgument)
	}

	for _, q := range p.Questions {
		if len([]rune(q.QuestionText)) < 2 {
			return fmt.Errorf("question text is too short: %w", core_errors.ErrInvalidArgument)
		}
		if len(q.Options) < 2 {
			return fmt.Errorf("question '%s' must have at least 2 options: %w", q.QuestionText, core_errors.ErrInvalidArgument)
		}

		for _, opt := range q.Options {
			if len([]rune(opt.OptionText)) == 0 {
				return fmt.Errorf("option text cannot be empty: %w", core_errors.ErrInvalidArgument)
			}
		}
	}

	return nil
}
