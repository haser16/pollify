package polls_postgres_repository

import core_postgres_pool "pollify/internal/core/repository/postgres/pool"

type PollsRepository struct {
	pool core_postgres_pool.Pool
}

func NewPollsRepository(pool core_postgres_pool.Pool) *PollsRepository {
	return &PollsRepository{
		pool: pool,
	}
}
