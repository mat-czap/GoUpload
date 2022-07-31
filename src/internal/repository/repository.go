package repository

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/matczap/repository/src/internal/folders"

	"context"
)

type FolderRepository struct {
	db    *dynamodb.Client
	table string
}

func NewFolderRepository(db *dynamodb.Client, table string) *FolderRepository {
	return &FolderRepository{db: db, table: table}
}

type folderItem struct {
	Pk      string         `dynamodbav:"pk"`
	Sk      string         `dynamodbav:"sk"`
	Payload folders.Folder `dynamodbav:"payload"`
}

func (r FolderRepository) SaveFolder(ctx context.Context, folder folders.Folder) error {
	folderItem := folderItem{
		Pk:      folder.Id,
		Sk:      folder.UserId,
		Payload: folder,
	}

	av, err := attributevalue.MarshalMap(folderItem)
	if err != nil {
		return err
	}

	_, err = r.db.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(r.table),
	})

	if err != nil {
		return err
	}

	return nil
}
