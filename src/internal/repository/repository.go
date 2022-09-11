package repository

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/matczap/repository/src/internal/folders"

	"context"
)

type FolderRepository struct {
	db    dynamoDBAPI
	table string
}

type dynamoDBAPI interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(options *dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}

func NewFolderRepository(db dynamoDBAPI, table string) *FolderRepository {
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

func (r FolderRepository) GetFolder(ctx context.Context, id string, userId string) (folders.Folder, error) {
	item, err := r.db.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String("folders"),
		Key: map[string]types.AttributeValue{
			"pk": &types.AttributeValueMemberS{Value: id},
			"sk": &types.AttributeValueMemberS{Value: userId},
		},
	})
	if err != nil {
		return folders.Folder{}, err
	}

	var folderItem folderItem

	err = attributevalue.UnmarshalMap(item.Item, &folderItem)
	if err != nil {
		return folders.Folder{}, err
	}

	return folderItem.Payload, nil
}
