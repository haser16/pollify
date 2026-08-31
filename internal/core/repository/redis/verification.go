package core_redis

import (
	"context"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type VerificationStore interface {
	Set(
		ctx context.Context,
		token string,
		userID int,
		ttl time.Duration,
	) error

	Get(
		ctx context.Context,
		token string,
	) (int64, error)

	Delete(
		ctx context.Context,
		token string,
	) error
}

type verificationStore struct {
	client *redis.Client
}

func NewVerificationStore(client *redis.Client) VerificationStore {
	return &verificationStore{
		client: client,
	}
}

func (s *verificationStore) Set(
	ctx context.Context,
	token string,
	userID int,
	ttl time.Duration,
) error {
	return s.client.Set(
		ctx,
		"verification:"+token,
		userID,
		ttl,
	).Err()
}

func (s *verificationStore) Get(
	ctx context.Context,
	token string,
) (int64, error) {
	value, err := s.client.Get(
		ctx,
		"verification:"+token,
	).Result()
	if err != nil {
		return 0, err
	}

	return strconv.ParseInt(value, 10, 64)
}

func (s *verificationStore) Delete(
	ctx context.Context,
	token string,
) error {
	return s.client.Del(
		ctx,
		"verification:"+token,
	).Err()
}
