package uploadprovider

import (
	"context"
	"food-client/modules/user/model"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*model.Avatar, error)
}
