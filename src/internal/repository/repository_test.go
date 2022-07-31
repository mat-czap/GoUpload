package repository_test

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/matczap/repository/src/internal/folders"
	"github.com/matczap/repository/src/internal/repository"
	"github.com/matczap/repository/src/pkg/dynamolocal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository(t *testing.T) {
	ctx := context.Background()
	t.Run("should save", func(t *testing.T) {
		client := dynamolocal.CreateLocalClient()
		result, err := client.ListTables(ctx, &dynamodb.ListTablesInput{})

		_ = result
		r := repository.NewFolderRepository(client, "folders")

		folder := folders.Folder{Id: "1", UserId: "1"}
		err = r.SaveFolder(ctx, folder)

		assert.NoError(t, err)

	})
}
