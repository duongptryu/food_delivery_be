package uploadprovider

import (
	"context"
	"food_delivery_be/common"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, adta []byte, dst string) (*common.Image, error)
}
