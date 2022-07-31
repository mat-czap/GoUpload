package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type folderGetter interface {
	GetFolder(ctx context.Context, id, userID string) (Folder, error)
}

func handler(ctx context.Context) (string, error) {

	return "hello", nil
}

func main() {
	lambda.Start(handler)
}
