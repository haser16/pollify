package polls_postgres_repository

import "time"

type PollModel struct {
	ID          int
	Title       string
	Description string
	CreatedAt   time.Time
	ExpiresAt   time.Time
	Completed   bool
	AuthorID    int
}
