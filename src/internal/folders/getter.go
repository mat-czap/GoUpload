package folders

import (
	"context"
	"time"
)

type FileGetter struct{}

type Folder struct {
	Id        string
	UserId    string
	Name      string
	CreatedAt time.Time
	Size      int
}

func (g FileGetter) GetFolder(ctx context.Context) {

}
