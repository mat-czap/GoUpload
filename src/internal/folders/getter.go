package folders

import (
	"context"
	"errors"
	"strings"
	"time"
)

type FolderGetter struct {
	repository folderRepository
}

type folderRepository interface {
	GetFolder(ctx context.Context, id, userID string) (Folder, error)
}

type Folder struct {
	Id        string
	UserId    string
	Name      string
	CreatedAt time.Time
	Size      int
}

func NewFolderGetter(r folderRepository) *FolderGetter {
	return &FolderGetter{repository: r}
}

func (g FolderGetter) GetFolder(ctx context.Context, id, userId string) (Folder, error) {
	if strings.HasPrefix(userId, "baned") {
		return Folder{}, errors.New("baned user")
	}

	return g.repository.GetFolder(ctx, id, userId)
}
