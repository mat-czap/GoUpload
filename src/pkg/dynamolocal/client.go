package dynamolocal

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func CreateLocalClient() *dynamodb.Client {
	return dynamodb.NewFromConfig(aws.Config{
		Region: "local",
		Credentials: aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     "local",
				SecretAccessKey: "local",
				SessionToken:    "local",
			}, nil
		}),
		EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) { // nolint
			dynamodbURL := "http://localhost:8000"

			return aws.Endpoint{URL: dynamodbURL}, nil
		}),
	})
}
