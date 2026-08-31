package core_publisher

import (
	"context"
)

type VerificationMessage struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type Publisher interface {
	Publish(ctx context.Context, message VerificationMessage) error
}

type Consumer interface {
	Consume(
		ctx context.Context,
		handler func(VerificationMessage) error,
	) error
}
