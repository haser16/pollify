package core_s3_aws

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"

	core_s3 "pollify/internal/core/repository/s3"
)

type Client struct {
	client *awsS3.Client
}

var _ core_s3.Storage = (*Client)(nil)

func NewClient(client *awsS3.Client) *Client {
	return &Client{
		client: client,
	}
}

func (c *Client) Upload(
	ctx context.Context,
	bucket string,
	key string,
	body io.Reader,
	contentType string,
) error {
	_, err := c.client.PutObject(
		ctx,
		&awsS3.PutObjectInput{
			Bucket:      aws.String(bucket),
			Key:         aws.String(key),
			Body:        body,
			ContentType: aws.String(contentType),
		},
	)

	return err
}

func (c *Client) Delete(
	ctx context.Context,
	bucket string,
	key string,
) error {
	_, err := c.client.DeleteObject(
		ctx,
		&awsS3.DeleteObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		},
	)

	return err
}
