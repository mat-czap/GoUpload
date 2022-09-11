package repository_test

import (
	"context"
	"testing"

	"github.com/matczap/repository/src/internal/folders"
	"github.com/matczap/repository/src/internal/repository"
	"github.com/matczap/repository/src/pkg/dynamolocal"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	ctx := context.Background()
	t.Run("should save", func(t *testing.T) {
		client := dynamolocal.CreateLocalClient()

		r := repository.NewFolderRepository(client, "folders")

		folder := folders.Folder{Id: "1", UserId: "1"}
		err := r.SaveFolder(ctx, folder)

		assert.NoError(t, err)

	})

	t.Run("should get", func(t *testing.T) {
		client := dynamolocal.CreateLocalClient()

		r := repository.NewFolderRepository(client, "folders")

		folderId := "1"
		userId := "1"

		item, err := r.GetFolder(ctx, folderId, userId)

		assert.NoError(t, err)
		assert.Equal(t, folders.Folder{Id: "1", UserId: "1"}, item)

	})
}
