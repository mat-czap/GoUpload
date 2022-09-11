package main

import (
	"os"

	"context"

	"github.com/matczap/repository/src/internal/folders"

	"github.com/matczap/repository/src/internal/repository"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type folderGetter interface {
	GetFolder(ctx context.Context, id, userID string) (folders.Folder, error)
}

type handlerType func(ctx context.Context) (folders.Folder, error)

func handlerConstructor(folderGetter folderGetter) handlerType {
	return func(ctx context.Context) (folders.Folder, error) {
		folder, err := folderGetter.GetFolder(ctx, "1", "123")
		return folder, err
	}
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	lambda.Start(handlerConstructor(folders.NewFolderGetter(repository.NewFolderRepository(svc, os.Getenv("FILE_TABLE")))))

}
