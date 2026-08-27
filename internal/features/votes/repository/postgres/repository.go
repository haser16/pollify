package votes_postgres_repository

import core_postgres_pool "pollify/internal/core/repository/postgres/pool"

type VotesRepository struct {
	pool core_postgres_pool.Pool
}

func NewVotesRepository(pool core_postgres_pool.Pool) *VotesRepository {
	return &VotesRepository{pool: pool}
}
