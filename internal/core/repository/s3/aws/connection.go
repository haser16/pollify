package core_s3_aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewSDKClient(
	ctx context.Context,
	endpoint string,
	region string,
	accessKey string,
	secretKey string,
) (*awsS3.Client, error) {
	awsConfig, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				accessKey,
				secretKey,
				"",
			),
		),
	)
	if err != nil {
		return nil, err
	}

	client := awsS3.NewFromConfig(
		awsConfig,
		func(options *awsS3.Options) {
			options.BaseEndpoint = aws.String(endpoint)
		},
	)

	return client, nil
}
