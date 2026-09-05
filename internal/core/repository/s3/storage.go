package core_s3

import (
	"context"
	"io"
)

type Storage interface {
	Upload(
		ctx context.Context,
		bucket string,
		key string,
		body io.Reader,
		contentType string,
	) error

	Delete(
		ctx context.Context,
		bucket string,
		key string,
	) error
}
